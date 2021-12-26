module.exports = {
    dest: "./dist",
    base: "/leafBot/dist/",
    title: 'LeafBot',
    description: '一个onebot的sdk',
    themeConfig: {
        sidebar: {
            "/data/quickstart/": [
                "env", "small_app"
            ]
        },
        nav: [
            {text: 'Home', link: '/'},                      // 根路径
            {text: "Quick Start", link: "/data/quickstart/env"},
            {text: "Driver", link: "/data/driver"},
            {text: "Config", link: "/data/config"},
            {text: "Github", link: "https://github.com/huoxue1/leafbot"},
            // 显示下拉列表
            // {
            //     text: 'Languages',
            //     items: [
            //         {text: 'Chinese', link: '/language/chinese'},
            //         {text: 'Japanese', link: '/language/japanese'}
            //     ]
            // },
        ]
    },
    markdown: {
        lineNumbers: true//代码显示行号
    }
}