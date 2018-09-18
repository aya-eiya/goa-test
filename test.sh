# /bin/sh

rm -rf apps
goagen app -d github.com/aya-eiya/goa-test/design/app -o apps/
goagen client -d github.com/aya-eiya/goa-test/design/app -o apps/
mv "apps/tool/-cli" "apps/tool/goa-test-cli"
go test -v ./apps/...
