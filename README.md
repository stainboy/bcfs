#TBD

    bcfs login [OPTIONS] EMAIL PASSWORD
    bcfs ls [OPTIONS] REMOTE_PATH
    bcfs mkdir [OPTIONS] REMOTE_PATH
    bcfs cp [OPTIONS] LOCAL_SOURCE... cfs:/REMOTE_DEST
    bcfs cp [OPTIONS] cfs:/REMOTE_SOURCE... LOCAL_DEST
    bcfs rm [OPTIONS] REMOTE_PATH
    bcfs mv [OPTIONS] REMOTE_SOURCE... REMOTE_DEST
    bcfs help [COMMAND]
    bcfs version

PS: When successfully logon, cookie is saved at ~/.bcfs.cookie. Remove the file to logout.

PS: When invoked as tcfs, cookie is named ~/.tcfs.cookie.