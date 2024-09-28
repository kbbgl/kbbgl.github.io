---
slug: 3-way-data-migration-between-support-systems
title: 3-Way Data Migration between Support Systems
description: Fill me up!
authors: [kbbgl]
tags: [bash,data engineering,jq,migration,pandas,python]
---

## Introduction

The company I work for decided a few months ago that we’ll be moving all customer tickets and assets from two separate systems (one for chat and one for old-school tickets) into a new, integrated system which provides both capabilites.
My task was to perform the migration between the systems.
Even though I’m not data engineer by any means, I accepted the challenge and thought it would teach me a lot about the planning and execution of such a complex project. It would also allow me to hone in my development/scripting skills and finally have some hands-on experience using a Python library I was always interested in working with, [pandas](https://pandas.pydata.org/).

## Planning the Migration

The first step was to understand what format the source and destination platforms supplied and expected the data in, respectively.
After reaching out to the destination platform engineers, they provided the following information:

- Supply a TAR archive with JSON files named `backup_tickets_{batch_number}.json`.
- The JSON files need to have the following structure:

    ```json
    {
       "data": {
           "tickets": {
              "data": [],
              "comments": [],
              "users": []
           }
       }
    }
    ```

- Each JSON file should not include more than 100 objects within each nested array (`data`, `comments`, `users`, `organizations`).
- The structure for each object had to be as follows:

  - `ticket` object blueprint:

    ```json
    {
       "created_at": "2021-01-01T00:00:00Z",
       "requester_id": 123456,
       "id": 654321
    }
    ```

  - `comment` object blueprint:

    ```json
    {
       "created_at": "2021-02-02T02:02:02Z",
       "ticket_id": 211243,
       "id": 789012,
       "public": true,
       "html_body": "<p>dummy message</p>",
       "author_id": 123456
    }
    ```

    - `user` object blueprint:

    ```json
    {
       "name": "John Doe",
       "id": 00076,
       "email": "john.doe@somedomain.org"
    }
    ```

The source data came from two different places:

- The chat-related data was stored in a database that I did not have access to. Therefore, I received the data in CSV format with the following headers:

  ```bash
  head -n1 chat_data.csv

  TICKET_CREATED_AT,TICKET_REQUESTER_ID,COMMENT_PART_ID,CONVERSATION_ID,COMMENT_PUBLIC,BODY,COMMENT_CREATED_AT,AUTHOR_ID,NAME,EMAIL
  ```

- The ticket-related data was available through the provider REST API. This data was already imported into the destination system when I was assigned this task. So I needed to focus my efforts on migrating the chat-related data. I thought it would make the process easier but it actually complicated it a bit. More details on that in the section below.

## The User Uniqueness Problem

One crucial detail that was not mentioned in the destination platform documentation was that the **users imported need to be unique based on their email address**.

This complicated things because:

- The users from the source ticketing platform already existed in the destination platform since they were previously imported.

- There were users that existed on both source platforms but had different id attribute values.
  
- The requirement from the destination platform system was to use the user id attribute that was already imported into the system.

After some brainstorming, I was able to think about how to solve this discrepency. I broke up the problem into the following steps:

1. Retrieve all users from the source ticketing platform.
1. Retrieve all users from the source chat platform.
1. Run a VLOOKUP to compare the id from both lists of users. If the email existed in both lists, take the id from the list of users generated in step 1. Otherwise, use the id from the list retrieved in step 2.

## Step 1: Retrieving Users from Source Ticketing Platform

Since I already knew the source ticketing platform could supply user data by accessing their REST API, I began by reading their documentation.

I found out that:

- The endpoint to retrieve user data was:
    `GET /api/v1/users.json`.

- The endpoint returned a maximum of 100 users in JSON format.
- The endpoint offered a query parameter to specify paging through all users, i.e. `GET /api/v1/users.json?page=100`.
- Basic authentication method was used to authenticate calls.
- The JSON response included many user attributes. I was only interested in 2 of them: `id` and `email`.

Reviewing the total amount of users in that platform, I found that there were a total of 56.7k. So I needed to run 567 calls (56700/100) to retrieve all the users. Sounds like a perfect solution for a combination of `curl`, `jq` within `bash` loops and redirection of output to appended file.
This is what I came up with:

```bash
for i in {1..567}
do
    echo "Retrieving users page $i..."
    curl https://$SOURCE_URL.com/api/v1/users.json\?page=$i -u my_user/my_password | jq '.users[] | [.id, .email] | @csv' >> ticket_platform_users.csv
    echo "Finished retrieving users from page $i"
done
```

This loop generated a CSV file with the following data:

```bash
head -n3 ticket_platform_users.csv
 
id,email
1,user1@somedomain1.org
2,user2@somedomain2.org
```

Step 1 is done!

## Step 2: Generating Unique Users List from Source Chat Platform

As previously mentioned, the data from the chat platform was sent to me in CSV format with the following headers:

```bash
  head -n1 chat_data.csv

  TICKET_CREATED_AT,TICKET_REQUESTER_ID,COMMENT_PART_ID,CONVERSATION_ID,COMMENT_PUBLIC,BODY,COMMENT_CREATED_AT,AUTHOR_ID,NAME,EMAIL
```

The `chat_data.csv` file included all chat conversations which meant that every line in the CSV represented one message sent on the chatting platform. Since users (or `AUTHOR_ID` in this case) can write multiple messages in different `CONVERSATION_ID`s, I needed a way to take only the unique `AUTHOR_ID`s from this CSV. This is where I was introduced to the power of the Python `pandas` library.
I was able to generate a list of unique users using the following script:

```python
#!/usr/bin/env python
"""
generate_unique_users.py
"""
import pandas as pd
 
# Read only the relevant columns from the CSV file
users = pd.read_csv(
   'chat_data.csv',
   usecols=['NAME', 'EMAIL', 'AUTHOR_ID']
)
 
# Rename the columns according to the `users` object blueprint
users_rename = users.rename({
   "AUTHOR_ID": "id",
   "NAME": "name",
   "EMAIL": "email"
})
 
# Modify the column data type according to `users` object blueprint. `int64` is the `numpy.dtype` representing a number
users_coltype = users_rename.astype({
   "id": "int64"
})
 
# Remove any rows with `null`
removed_na = users_coltype.dropna(how='any')
 
# Remove duplicate rows, use `email` column as key
removed_dups = removed_na.drop_duplicates(subset=["email"], keep="first")
 
# Display the number of unique users
count = removed_dups.index
print("Found {} unique users".format(len(count)))
 
# Generate new CSV with unique users
removed_dups.to_csv("chat_platform_users.csv", index=False)
```

I ran the script which generated `chat_platform_users.csv` with 13k unique users (where the file `chat_data.csv` had 2.5M rows).

Step 2 done!

## Step 3: User List Merging based on Common Attribute

So now we had 2 CSV files:

- `ticket_platform_users.csv` with 56k users.
- `chat_platform_users.csv` with 13k users.

Since I am pretty proficient with Excel/Google Spreadsheet formulas, I decided to use the `VLOOKUP` function to generate the required list.
I imported both CSVs into two separate sheets named after their filenames, namely:

- Sheet 1: `ticket_platform_users` had 2 columns, `email, id`
- Sheet 2: `chat_platform_users`, had 3 columns, `name, email, id`.
- Sheet 3: `merged` – This is where we’re going to generate a merged list of users. It had 3 columns as well, `name, email, id`.

The name and email rows were taken from Sheet 2, the `id` had the following formula in it (in the 2nd row of the sheet):

```excel-formula
=IFNA(VLOOKUP(A2, ticket_platform_users!$A$2:$B$56748, 2, false), chat_platform_users!C2)
```

The explanation of the formula:

- The `IFNA` function takes in 2 arguments:

  1. The first argument is the value which we want to check if it returns a N/A value.
  
  1. The second argument determines what value will be returned in the cell if the value returned in the first argument is NA.

  We use this function because we know that some user emails (cells `A2...AN`) will not be found in both sheets so we want to return the `id` of the user from the `chat_platform_users` in this case. Otherwise, if the `VLOOKUP` function does return a successful match for email in both sheets, the `IFNA` will evaluate the `VLOOKUP` value instead.

- The `VLOOKUP` function takes 4 arguments:

    1. The first argument is the value to search for. In the example above, cell `A2` evaluates to a user email so we’ll be searching for that email.

    1. The second argument is the range where we’ll be looking for a value matching `A2`. `ticket_platform_users!$A$2:$B$56748` in plain words means: ‘within the ticket_platform_users spreadsheet in the static range from cells `A2` to `B56748`‘.

    1. The third argument identifies the column number to return if a successful match is found. We’ve specified 2 in this case because we’re interested in returning the value from 2nd column from the sheet `ticket_platform_users` which is the id.

    1. The fourth argument is optional and it indicates to the function whether the data is sorted. It is set to false because I imported the data unsorted.

I exported the merged sheet to file `merged_users.csv`.

Step 3 done!

## Creating Batch JSON Objects from CSV

At this point, I had put a lot of time and mental energy into manipulating the users data to fit the requirements but had not touched either the tickets or comments data yet. I still needed to somehow generate the JSON representation for all 3 data sets (`user`, `comment`, `ticket`) and somehow insert these data sets into their appropriate final JSON object arrays.

Luckily, `pandas` was there to lend a hand.

I wrote 3 similar scripts, one for each data set.
For the `user` CSV:

```python
#!/usr/bin/env python
"""
# generate_user_json.py
"""

import pandas as pd
 
users = pd.read_csv(
    'merged_users.csv'
)
 
# Modify column data types so `id` is a number
users_dt = users.astype(
    {
        'id': 'int64'
    }
)
 
 
# Convert DataFrame to JSON object
# One object per line
out = users_dt.to_json(orient='records', lines=True)
 
with open('all_users_lines.json', 'w', encoding='utf-8') as f:
    f.write(out)
"""
```

For `comment`:

```python
#!/usr/bin/env python
"""
generate_comment_json.py
"""

import pandas as pd
 
comments = pd.read_csv(
    'chat_data.csv',
     usecols=[
        'COMMENT_PART_ID', 
        'AUTHOR_ID', 
        'BODY', 
        'COMMENT_PUBLIC', 
        'CONVERSATION_ID', 
        'COMMENT_CREATED_AT'
    ]
)
 
# Remove duplicates
removed_dups = comments.drop_duplicates(subset=["COMMENT_PART_ID"], keep="first")
 
# Remove comments with null fields
removed_na = removed_dups.dropna(how='any')
 
# Rename columns according to specifications
comments_object = removed_na.rename(columns=
    {
        "COMMENT_PART_ID": "id", 
        "AUTHOR_ID": "author_id", 
        "BODY": "html_body",
        "COMMENT_PUBLIC": "public",
        "COMMENT_CREATED_AT": "created_at",
        "CONVERSATION_ID": "ticket_id"
    }
)
 
# Modify column data types
# According to numpy dtypes
# https://numpy.org/doc/stable/user/basics.types.html
comments_final = comments_object.astype(
    {
        'id': 'int64', 
        'author_id': 'int64',
        'html_body': 'str',
        'public': 'bool',
        'ticket_id': 'int64',
        'created_at': 'str'
    }
)
 
# Convert DataFrame to JSON object
# One object per line
out = comments_final.to_json(orient='records', lines=True)
outfile='all_comments_lines.json'
with open(outfile, 'w', encoding='utf-8') as f:
    f.write(out)
```

For `ticket`:

```python
#!/usr/bin/env python

""" 
generate_ticket_json.py
"""

import pandas as pd
 
tickets = pd.read_csv(
    'chat_data.csv', 
    usecols=['TICKET_REQUESTER_ID', 'TICKET_CREATED_AT', 'CONVERSATION_ID']
)
 
# Remove ticket duplicates
removed_ticket_dups = tickets.drop_duplicates(subset=["CONVERSATION_ID"], keep="first")
 
# Rename columns according to specifications
tickets_object = removed_ticket_dups.rename(columns=
    {
        "CONVERSATION_ID": "id", 
        "TICKET_CREATED_AT": "created_at", 
        "TICKET_REQUESTER_ID": "requester_id"
    }
)
 
# Modify column data types
tickets_final = tickets_object.astype(
    {
        'id': 'int64', 
        'requester_id': 'int64'
    }
)
 
out = tickets_final.to_json(orient='records', lines=True)
outfile='all_tickets_lines.json'
with open(outfile, 'w', encoding='utf-8') as f:
    f.write(out)
```

I used the `DataFrame::to_json(lines=True)` method, option because it generated a JSON object per line. Since I knew that the requirement was to have a maximum of 100 objects per array representing the comment, user or ticket in the final JSON asset, I found it easier to split the output files (`all_tickets_lines.json`, `all_comments_lines.json`, `all_users_lines.json`) into a batch of files where each one would have 100 JSON objects within them.

I developed the following script to produce just that:

```python
#!/usr/bin/env python
"""
split_into_batch_files.py
"""
# Define iteration group
assets=['tickets', 'users', 'comments']
 
 
"""
Input JSON files
"""
tickets_file="../../{}/all_{}_lines.json".format(assets[0], assets[0])
users_file="../../{}/all_{}_lines.json".format(assets[1], assets[1])
comments_file="../../{}/all_{}_lines.json".format(assets[2], assets[2])
 
 
### Split JSON files into smaller batches
lines_per_file=100
 
for asset in assets:
    current_file='../../{}/all_{}_lines.json'.format(asset, asset)
    print("Splitting file '{}' into files with {} lines...".format(current_file, str(lines_per_file)))
    batch_file = None
    with open(current_file) as bigfile:
        for line_number, line in enumerate(bigfile):
            if line_number % lines_per_file == 0:
                if batch_file:
                    batch_file.close()
                small_filename = './{}_split/{}.json'.format(asset, line_number + lines_per_file)
                batch_file = open(small_filename, "w")
            batch_file.write(line)
        if batch_file:
            batch_file.close()
    print("Finished splitting file '{}'".format(current_file))
```

After executing the script above, I had files with 100 JSON objects on each line in the file:

```bash
ls ./comments_split/*.json | wc -l
   21004
 
ls ./tickets_split/*.json | wc -l
     925
 
ls ./users_split/*.json | wc -l
     138
```

The final step was to iterate over each folder, read the `n`th file within that folder and insert the 100 JSON objects into an output JSON file named `backup_tickets_{batch_number}.json`.
This is the script I wrote to produce just that:

```python
#!/usr/bin/env python

import os
import pandas as pd
import json
 
assets=['tickets', 'users', 'comments']
 
"""
Input JSON files
"""
tickets_dir="./{}_split".format(assets[0])
comments_dir="./{}_split".format(assets[2])
users_dir="./{}_split".format(assets[1])
 
"""
Output JSON files
"""
output_folder="../assets_to_send"
 
# Every file in asset folder is called after the batch number * 100
# So first file is called '100.json' per each folder ('./tickets/100.json', './users/100.json', './comments/100.json')
# The last file is called './comments_split/2100400.json
for x in range(100, 2100400, 100):
 
    output_filename=os.path.join(output_folder, "backup_tickets_{}.json".format(x))
 
    print("Generating JSON for batch {}...".format(x))
 
    # Create root JSON object
    root={}
    root["data"] = {}
    root["data"]["tickets"] = {}
    root["data"]["tickets"]["data"] = []
    root["data"]["tickets"]["comments"] = []
    root["data"]["tickets"]["users"] = []
    root["data"]["tickets"]["organizations"] = []
 
    for asset in assets:
         
        current_ticket_file = os.path.join(tickets_dir, "{}.json".format(x))
        current_comments_file = os.path.join(comments_dir, "{}.json".format(x))
        current_users_file = os.path.join(users_dir, "{}.json".format(x))
 
        if os.path.exists(current_ticket_file):
            tickets=pd.read_json(current_ticket_file, lines=True, convert_dates=False)
            ticket_out = tickets.to_json(orient='records')
            ticket_out_dict = json.loads(ticket_out)
            root["data"]["tickets"].update({'data': ticket_out_dict})            
         
        if os.path.exists(current_users_file):
            users = pd.read_json(current_users_file, lines=True, convert_dates=False)
            users_out = users.to_json(orient='records')
            users_out_dict = json.loads(users_out)
            root["data"]["tickets"].update({'users': users_out_dict})
 
        if os.path.exists(current_comments_file):
            comments = pd.read_json(current_comments_file, lines=True, convert_dates=False)
            comments_out = comments.to_json(orient='records')
            comments_out_dict = json.loads(comments_out)
            root["data"]["tickets"].update({'comments': comments_out_dict})
 
    print("Writing ticket JSON to file '{}'...".format(output_filename))
    with open(output_filename, 'w', encoding='utf-8') as f:
        json.dump(root, f, indent=3)
```

I compressed the generated assets into an archive which contained 21k JSON files:

```bash
tar tf backup_tickets.tar.gz | head
./backup_tickets_1605900.json
./backup_tickets_1670800.json
./backup_tickets_1585100.json
./backup_tickets_1217200.json
./backup_tickets_1344500.json
./backup_tickets_990600.json
./backup_tickets_1331400.json
./backup_tickets_1262300.json
./backup_tickets_171600.json
 
tar tf backup_tickets.tar.gz| wc -l
   21004
```

Migration complete!
