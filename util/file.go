package util


import (
    "path"
)

//发送HTTP download
func Extension(name string) string {
    ext := path.Ext(name)
    if ext != "" {
        return ext[1:]
    }

    return ext
    //slices := strings.Split(name, ".")
    //dots := len(slices)
    //
    //if dots == 1 {
    //    return ""
    //}
    //
    //i := 1
    //if dots > 3 {
    //    i = dots - 2
    //}
    //return strings.Join(slices[i:], ".")
}




