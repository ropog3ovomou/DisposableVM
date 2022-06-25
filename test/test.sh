cd "$(dirname "$0")"
trap on_exit EXIT
#set -x
set -e
vm=appDir
orig_dir="`pwd`"
on_exit() {
	cd "$orig_dir"
	rm -r $vm
	if test -z $success; then
		echo === FAIL ===
	else
		echo === SUCCESS ===
	fi
}
mkdir $vm
cp ../disposable $vm
cp -r ../bin $vm
cp ../win.conf $vm/win.conf.orig
cp ../mac.conf $vm/mac.conf.orig
cp *.conf $vm
pushd $vm >/dev/null
./disposable configure ../fakeVM/*.vmx
./disposable start
./disposable dispose
popd >/dev/null
success=1
