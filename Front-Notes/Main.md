

# 教学路径

Vite脚手架快速搭建简单的单页面Vue实例模板

​	=> 1.Type Script基础语法

​         2.Vue基础语法

​	在单页面中基本演示上述的知识点，知识点见菜鸟教程的大纲



二、从单页面过渡到多页面

​	=> 1.Vue-Router路由引入，介绍它是如何单页面模拟切换多页面的，重点讲baseURL

​		 2.组件的引入，在多个页面中会有重复出现的功能模块，通过编写组件，由组件组成页面

​		 3.组件间的通信，父传子，子传父

​		 4.Pinia全局状态管理引入，相当于全局变量，在多个页面间传输数据



三、前后端分离

​	=> 1.Axios引入，介绍前后端通信，Axios的相关语法，Axios的请求的二次封装作为进阶，让他们自行探索，大作业 写得好有加分

​		 2.Apifox的高级Mock的使用，和介绍，讲出其实前后端分离就是前端Vite开出来一台服务器，后端也开出来一台服务器，二者开放相同的端口和api进行通信，达到前后端分离的效果。Apifox的高级Mock相当于也开了一台服务器，使得前端服务器和后端服务器能相互发请求….

​		 3.讲讲跨域还有Http请求的原理

 





# 0 前置

## 0.1包管理工具

### yarn

```shell
npm i -g yarn --registry=https://registry.npmmirror.com
```



### pnpm（推荐）

```shell
npm install -g pnpm
```



## 0.2Vite脚手架快速搭建Vue项目实例

```shell
npm init vite@latest my-vue-app -- --template vue
```

`my-vue-app`是创建的项目的名称

`temple vue`是指快速搭建一个Vue实例模板



## 0.3Eslint配置

### 0.3.1 Eslint初始化配置

```shell
npm i -g eslint

eslint --init
```



### 0.3.2 Vscode插件安装

![image-20230622165620204](Main/image-20230622165620204.png)



### 0.3.3 项目根目录配置.eslintrc.cjs

```typescript
module.exports = {
    env: {
       browser: true,
       es2021: true,
       node: true,
     },
     extends: [
       "plugin:vue/vue3-essential",
       "@vue/eslint-config-typescript/recommended"
     ],
     overrides: [
     ],
     parserOptions: {
       ecmaVersion: "latest",
       sourceType: "module",
       project: "./tsconfig.json"
     },
     plugins: [
       "vue"
     ],
     rules: {
       quotes: ["warn", "double"],
       semi: "error",
       "vue/multi-word-component-names": "off",
       "@typescript-eslint/ban-ts-comment": "off",
       "no-trailing-spaces": "error"
     }
   };
```



---

# 1 TypeScript 基础语法

参考[入门教程 - TypeScript 中文文档 (nodejs.cn)](https://nodejs.cn/typescript/handbook/)

[TypeScript 教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/typescript/ts-tutorial.html)



---

# 2 Vue3 

参考 [Vue.js - 渐进式 JavaScript 框架 | Vue.js (vuejs.org)](https://cn.vuejs.org/)

[Vue3 教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/vue3/vue3-tutorial.html)



### 2.0 引言

一个前端项目的开发思路是，将一个项目的业务根据功能拆分成多个**页面（page）**，当前流行的**单页面（SPA）**应用，采用**VueRouter路由**的方式实现在单页面中**模拟跳转页面**。

而每个页面中，根据功能需求，将页面拆分成多个**组件（Component）**，以提高代码的复用度，举个简单的例子，在一个项目中可能有多个业务场景会用到表单（form），输入框（input），警示框（alert），只需要写一个组件即可在全项目中通用，当然组件与**组件的通信和数据传输**是后话，在组件2.3中会教到。组件和组件甚至还可以嵌套，一个大的组件里面可能有多个小组件构成，这些都是需要程序编写前整体考虑构思到的。



### 2.1 响应式

2.2  

### 2.3 组件

2.3 

（这里还只在单页面中把大部分的语法举例子）



### 2.4 VueRouter
