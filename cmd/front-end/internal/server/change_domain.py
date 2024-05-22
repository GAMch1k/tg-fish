import os

dirs = [
    './static',
    ]
REPLACE_IT = 'http://localhost:3030'
TO_IT = 'http://localhost:3030'


for rootdir in dirs:
    for subdir, dirs, files in os.walk(rootdir):
        for file in files:
            print(os.path.join(subdir, file))
            if ('.html' in file or '.htm' in file):
                data = ''
                with open(os.path.join(subdir, file), 'r', encoding='utf8') as f:
                    data = f.read()
                data2 = data.replace(REPLACE_IT, TO_IT)
                with open(os.path.join(subdir, file), 'w', encoding='utf8') as f:
                    f.write(data2)
