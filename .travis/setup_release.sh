grep -rl 'import "../Controls"' ./src/ui | xargs sed -i -e 's/import "..\/Controls"/\/\/import "..\/Controls"/g'
grep -rl 'import "../"' ./src/ui | xargs sed -i -e 's/import "..\/"/\/\/import "..\/"/g'
grep -rl 'import "Delegates/"' ./src/ui | xargs sed -i -e 's/import "Delegates\/"/\/\/import "Delegates\/"/g'
grep -rl 'import "../Dialogs/"' ./src/ui | xargs sed -i -e 's/import "..\/Dialogs\/"/\/\/import "..\/Dialogs\/"/g'
grep -rl 'import "Dialogs/"' ./src/ui | xargs sed -i -e 's/import "Dialogs\/"/\/\/import "Dialogs\/"/g'
grep -rl 'import "Controls"' ./src/ui | xargs sed -i -e 's/import "Controls"/\/\/import "Controls"/g'
grep -rl 'import "../Utils"' ./src/ui | xargs sed -i -e 's/import "..\/Utils"/\/\/import "..\/Utils"/g'
grep -rl '// url := core.NewQUrl3("qrc' main.go | xargs sed -i -e 's/\/\/ url := core.NewQUrl3("qrc/url := core.NewQUrl3("qrc/g'
grep -rl 'url := core.NewQUrl3("src' main.go | xargs sed -i -e 's/url := core.NewQUrl3("src/\/\/url := core.NewQUrl3("src/g'
grep -rl '// import "qrc:/ui/src/ui' ./src/ui | xargs sed -i -e 's/\/\/ import "qrc:\/ui\/src\/ui/import "qrc:\/ui\/src\/ui/g'
find . -path "*-e" -delete
