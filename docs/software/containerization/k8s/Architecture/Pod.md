## `Pod`

Smallest unit of orchestration. The `Pod` can include many containers (`initContainer`s as well), has one IP address.


### `containers`

Each `Pod` has a number of containers. We can manage the resources used by the containers using the `PodSpec`:

### `initContainer`s

Containers that must finish running before the `container`s are run.