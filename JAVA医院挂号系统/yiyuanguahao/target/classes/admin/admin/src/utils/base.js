const base = {
    get() {
        return {
            url : "http://localhost:8080/yiyuanguahao/",
            name: "yiyuanguahao",
            // 退出到首页链接
            indexUrl: 'http://localhost:8080/yiyuanguahao/front/index.html'
        };
    },
    getProjectName(){
        return {
            projectName: "医院挂号系统"
        } 
    }
}
export default base
