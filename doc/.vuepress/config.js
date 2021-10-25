module.exports = {
    base: "/leafBot/",
    title: 'LeafBot',
    description: '一个onebot的sdk',
    themeConfig: {
        sidebar: 'auto',
        nav: [
            {text: 'Home', link: '/'},                      // 根路径
            {text: "Quick Start",link: "/data/quick_start"},
            {text: "Driver",link: "/data/driver"},
            {text: "Config",link: "/data/config"},
            {text: "Github",link: "https://github.com/huoxue1/leafBot"},
            // 显示下拉列表
            {
                text: 'Languages',
                items: [
                    {text: 'Chinese', link: '/language/chinese'},
                    {text: 'Japanese', link: '/language/japanese'}
                ]
            },
        ]
    }
                }