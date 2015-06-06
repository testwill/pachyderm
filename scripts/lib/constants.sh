# constants

BTRFS_MAJOR_VERSION="${BTRFS_MAJOR_VERSION:="3"}"
BTRFS_MINOR_VERSION="${BTRFS_MINOR_VERSION:="14"}"
DOCKER_MAJOR_VERSION="${DOCKER_MAJOR_VERSION:="1"}"
DOCKER_MINOR_VERSION="${DOCKER_MINOR_VERSION:="5"}"
PFS_DIR="${PFS_DIR:="${HOME}"}"
PFS_VOL="${PFS_VOL:="${PFS_DIR}/vol"}"
PFS_DATA_IMG="${PFS_DATA_IMG:="${PFS_DIR}/data.img"}"
PFS_CONTAINER_NAME="${PFS_CONTAINER_NAME:="pfs"}"
PFS_IMAGE="${PFS_IMAGE:="pachyderm/pfs"}"
PFS_PORT="${PFS_PORT:="650"}"
# TODO: is IMAGE_BUCKET legacy?
PFS_S3_BUCKET=${IMAGE_BUCKET:="pachyderm-images"}
