#include <iostream>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <dirent.h>
#include <fstream>
#include <set>

struct dirent *readdir (DIR *dirp);

int lstat(const char *path, struct stat *buf);
int closedir(DIR *dirp);


using namespace std;

int main(int argc, char *argv[]) {
    const char *path = argv[1];
    set<string> answer;
    DIR *directory = opendir(path);
    auto current = readdir(directory);
    while(current != NULL) {
        if(current->d_name[0] == '.') {
            current = readdir(directory);
            continue;
        }
        string d_name = current->d_name;
        int length = d_name.size();
        for(int i = length - 1; i >= 0; i--) {
            if (d_name[i] == '.') {
                if (d_name.substr(i + 1) == "html") {
                    ifstream f(string(argv[1]) + "/" + d_name);
                    string line;
                    while (getline(f, line)) {
                        unsigned long long j = 0;
                        bool check = false;
                        string str;
                        while(j < line.size()) {
                            if(line[j] == '<') {
                                check = true;
                                j++;
                                if(line[j] == '/') j++;
                                continue;
                            } else {
                                if(line[j] == '>' || line[j] == ' ') {
                                    check = false;
                                    answer.insert(str);
                                    str = "";
                                } else {
                                    if(!((line[j] >= 'A' && line[j] <= 'Z') || (line[j] >= 'a' && line[j] <= 'z'))) {
                                        str = "";
                                        break;
                                    }
                                    if(check) str += line[j];
                                }
                            }
                            j++;
                        }
                    }
                }
            }
        }
        cout << current->d_name << endl;
        current= readdir(directory);
    }
    ofstream fout("output.txt");
    for(auto i : answer) {
        fout << i << endl;
    }
    closedir(directory);
    fout.close();
    return 0;
}
