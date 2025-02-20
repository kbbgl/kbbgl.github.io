---
title: Splunk Cheat Sheet
slug: splunk-cheat-sheet
author: kgal-akl
tags: [devops, observability, tools, logs, audit, cheatsheet]
---

# Splunk Search Processing Language (SPL) Cheat Sheet

## 🔍 Basic Search

### Searches for keyword

```splunk-spl
index=my_index keyword
```

### Filters events where field matches value

```splunk-spl
index=my_index field=value
```

### Searches the last 7 days

```splunk-spl
index=my_index earliest=-7d latest=now
```

### Returns the first 10 results.
```splunk-spl
index=my_index | head 10
```

## 📌 Field Selection & Formatting

### Displays only selected fields in a table
```splunk-spl
index=my_index | table field1, field2, field3
```

### Keeps only field1 and field2 in results

```splunk-spl
index=my_index | fields field1, field2
```

### Removes field1 from results

```splunk-spl
index=my_index | fields - field1
```

## 📊 Sorting

### Sorts results in ascending order of `_time`

```splunk-spl
index=my_index | sort _time
```

### Sorts in descending order of count

```splunk-spl
index=my_index | sort - count
```

## 📈 Stats & Aggregation

### Counts occurrences of each field value

```splunk-spl
index=my_index | stats count by field
```

### Calculates the average of field

```splunk-spl
index=my_index | stats avg(field) as average
```

### Finds the first and last occurrence timestamps

```splunk-spl
index=my_index | stats min(_time) as first_seen, max(_time) as last_seen
```

### Computes the sum of field

```splunk-spl
index=my_index | stats sum(field) as total
```

## 🔄 Transforming Data

### Renames a field

```splunk-spl
index=my_index | rename field AS new_field
```

### Creates a new field by adding `field1` and `field2`

```splunk-spl
index=my_index | eval new_field=field1 + field2
```

### Creates a conditional field

```splunk-spl
index=my_index | eval status=if(response_code=200, "Success", "Fail")
```

## ⏳ Time Functions

### Converts epoch `_time` to human-readable format

```splunk-spl
index=my_index | convert ctime(_time)
```

### Formats _time

```splunk-spl
index=my_index | eval new_time=strftime(_time, "%Y-%m-%d %H:%M:%S")
```

### Searches the last 30 days, rounding to the start of the day

```splunk-spl
index=my_index earliest=-30d@d latest=now
```

## 🔄 Filtering & Matching

### Matches field with multiple values

```splunk-spl
index=my_index field IN ("value1", "value2")
```

### Excludes value from field

```splunk-spl
index=my_index field!=value
```

### Filters field for regex match

```splunk-spl
index=my_index | regex field="error|fail"
```
