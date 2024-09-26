## Volume Types

There are several types of volumes. Some local, some network.

in GCE/AWS, you can use volumes of types `GCEpersistentDisk` or `awsElasticBlockStore` which mount `GCE` or `EBS` disks in the `Pod`s.

`emptyDir` and `hostPath` are easy-to-use volumes. 

`hostPath` mounts a resource from the host node filesystem. There are two types which create resources on the host and use them if they don't exist already:

- `DirectoryOrCreate`
- `FileOrCreate`

