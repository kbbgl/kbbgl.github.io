---
title: Ansible Cheat Sheet
---

### `ansible` relevant files

* `/etc/ansible/hosts` - the list of remote servers that we'll manage.

 We can configure hosts and credentials:

    ```yaml
    [my_group]
    10.50.20.21
    10.50.20.22

 [my_group:vars]
    ansible_user=$myuser
    ansible_password=$mypassword
    ```

* `/etc/ansible/ansible.cfg` - `ansible` configuration.

### Commands

#### Ping

```bash
ansible my_group -m ping
```

#### Run specific command

```bash
ansible my_group -a "cat /etc/os-release"
```

### Run a Playbook

```bash
ansible-playbook my_playbook.yaml
```

### Run a Playbook with a variable

```bash
ansible-playbook my_playbook.yaml -e "package=nano"
```

### Playbook

Playbooks contains Plays which contain Tasks.

Playbooks are written in`YAML`.

A playbook example:

```yaml
--- 
  - name: check_nano # Start of Play
    hosts: my_group
    tasks: 
      - name: ensure nano is there # Start of task
        apt: # module
          name: nano
          state: latest
```

A playbook with variables:

```yaml
--- 
  - name: check_nano 
    hosts: my_group
    vars:
      package: nano
    tasks: 
      - name: ensure {{ package }} is there 
        apt: 
          name: "{{ package }}"
          state: latest
```

We can run this playbook with a variable as an argument:

```bash
ansible-playbook my_playbook.yml -e "package=httpd"
```

A Playbook with a loop:

```yaml
--- 
  - name: check_nano
    hosts: my_group
    vars:
      packages:
        - nano
        - vi
    tasks: 
      - name: install packages
        apt: 
          name: "{{ item }}"
          state: installed
        loop: "{{ packages }}" # loop
```
