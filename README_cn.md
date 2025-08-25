<p align="center">
  <a href="https://veops.cn/oneterm">
    <img src="https://github.com/user-attachments/assets/ab00344b-462b-44b9-9113-9fe735dfb096" />
  </a>
</p>

<p align="center">
  <a href="https://github.com/veops/oneterm/blob/main/LICENSE"><img src="https://img.shields.io/github/license/veops/oneterm" alt="Apache License 2.0"></a>
  <a href="https://github.com/veops/oneterm/releases"><img alt="the latest release version" src="https://img.shields.io/github/v/release/veops/oneterm?color=75C1C4&include_prereleases&label=Release&logo=github&logoColor=white"></a>
  <a href=""><img src="https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c" alt="go>=1.18"></a>
  <a href="https:https://github.com/sendya/ant-design-pro-vue"><img src="https://img.shields.io/badge/UI-Ant%20Design%20Pro%20Vue-brightgreen" alt="UI"></a>
  <a href="https://github.com/veops/oneterm/stargazers"><img src="https://img.shields.io/github/stars/veops/oneterm" alt="Stars Badge"/></a>
  <a href="https://github.com/veops/oneterm"><img src="https://img.shields.io/github/forks/veops/oneterm" alt="Forks Badge"/></a>
</p>

<h4 align="center">
 一款简单、轻量、灵活的堡垒机服务.
</h4>

<p align="center">
  <a href="https://trendshift.io/repositories/8690" target="_blank"><img src="https://trendshift.io/api/badge/repositories/8690" alt="veops%2Foneterm | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>
</p>

<p align="center">
  <a href="README.md">English</a> · 中文(简体)
</p>

## OneTerm是什么

OneTerm 是一款简单、轻量、灵活的企业级堡垒机产品，基于4A理念，即认证(Authen)、授权(Authorize)、账号(Account)、审计(Audit)设计开发，通过严格的访问控制和监控功能，确保系统的安全性和合规性。

- 官网: [veops.cn/oneterm](https://veops.cn/oneterm)
- 产品文档：[veops.cn/docs/docs/oneterm/onterm_design](https://veops.cn/docs/docs/oneterm/onterm_design)
- 在线体验：[OneTerm](https://term.veops.cn/oneterm/workstation)
  - 用户名: demo 或者 admin
  - 密码: 123456
- **注意**：`main` 分支在开发过程中可能处于**不稳定的状态**。请通过 [releases](https://github.com/veops/oneterm/releases) 获取最新稳定版本。

## 核心功能

+ **访问控制**： OneTerm作为一个中介站点，限制了对关键系统的直接访问。用户必须首先通过OneTerm进行身份验证，然后才能访问其他服务器或系统。

+ **安全审计**： OneTerm可以记录用户的登录和活动，提供审计日志，以便在发生安全事件时进行调查。这有助于确保每个用户的行为都是可追踪和可审计的。

+ **跳板访问**： OneTerm提供了一种跳板的方式，用户可以通过OneTerm连接到其他内部服务器。这种方式有助于减少直接暴露内部服务器的风险，因为只有OneTerm需要对外可访问。

+ **密码管理**： OneTerm可以实施强化的密码策略，并通过单一入口点集中管理密码。这有助于提高整个系统的密码安全性。

+ **会话录制**： OneTerm可以录制用户与服务器的会话，这对于监控和调查特权用户的活动非常有用。如果发生安全事件，可以回放会话录制以了解详细的操作。

+ **防止直接攻击**： 由于OneTerm是系统和资源的唯一入口，它可以成为攻击者攻击的主要障碍。这有助于减少直接攻击内部系统的风险。

+ **统一访问**： OneTerm提供了一个单一的入口点，用户可以通过它访问不同的系统，而无需记住多个登录凭据。这提高了用户的便利性和工作效率。

## 产品优势

+ **身份验证与授权**： Oneterm应具备强大灵活的的身份验证和授权机制。这包括支持多因素身份验证，确保只有授权用户能够访问内部网络资源，并能够细粒度地管理用户权限。

+ **安全通信**： Oneterm支持安全的通信协议和加密技术，以保护用户与内部服务器之间的数据传输。这有助于防止中间人攻击和数据泄露。

+ **审计和监控**： Oneterm具备强大的审计和监控功能，记录用户活动并生成审计日志。这有助于追踪安全事件、识别潜在的威胁，并满足合规性需求。

+ **远程管理和会话隔离**： Oneterm支持远程管理，使管理员能够安全地管理内部服务器。同时，它应该具备会话隔离功能，确保用户之间的访问是相互隔离的，防止横向扩展攻击。

+ **与开源CMDB紧密结合**：Oneterm与[维易CMDB](https://github.com/veops/cmdb)（已开源）紧密结合，用户可一键导入CMDB中的资产，确保操作方便，流程通畅。

## 技术栈

+ 后端：Go
+ 前端：Vue.js
+ UI组件库：Ant Design Vue

## 关注我们

欢迎 Star 加关注，第一时间获取更新动态！

![star us](https://github.com/user-attachments/assets/75c03659-4200-469e-b210-087a4d4473b6)

## 项目概览

<table>
  <tr>
    <td style="padding: 5px;background-color:#fff;">
      <img width="400" alt="dashboard" src="https://github.com/user-attachments/assets/cfbb7ae9-ddd3-4f0f-a37b-18b68bd8c7ac" />
    </td>
    <td style="padding: 5px;background-color:#fff;">
      <img width="400" alt="terminal" src="https://github.com/user-attachments/assets/e37f0ce8-d07c-42e0-a603-028b75c8e932" />
    </td>
  </tr>

  <tr>
    <td style="padding: 5px;background-color:#fff;">
      <img width="400" alt="work_station" src="https://github.com/user-attachments/assets/48a11f13-88be-4ec1-aa06-fbaade5721b1" />
    </td>
    <td style="padding: 5px;background-color:#fff;">
      <img width="400" alt="access_auth" src="https://github.com/user-attachments/assets/d1db6c5f-3ac0-46a1-9464-34c4c59243ed" />
    </td>
  </tr>

  <tr>
    <td style="padding: 5px;background-color:#fff;">
      <img width="400" alt="system_settings" src="https://github.com/user-attachments/assets/b9948d82-071a-427b-884f-a69fd37b27ae" />
    </td>
    <td style="padding: 5px;background-color:#fff;">
      <img width="400" alt="access_time" src="https://github.com/user-attachments/assets/55298679-8ded-4948-9738-38e418c2d03d" />
    </td>
  </tr>
</table>

## 开箱即用

### 方式一：快速部署（默认密码）
+ docker-compose 安装
  ```bash
  git clone https://github.com/veops/oneterm.git
  cd oneterm/deploy
  docker compose up -d
  ```

### 方式二：安全部署（自定义密码）
+ 生产环境建议使用安装脚本配置安全密码：
  ```bash
  git clone https://github.com/veops/oneterm.git
  cd oneterm/deploy
  ./setup.sh
  docker compose up -d
  ```
  安装脚本将会：
  - 生成安全的随机密码或允许你设置自定义密码
  - 使用你的密码更新所有配置文件
  - 为安全起见创建备份文件

+ **访问**
  - 浏览器: [http://127.0.0.1:8666](http://127.0.0.1:8666)
  - 账号: admin
  - 密码: 123456（默认）或使用setup.sh时设置的自定义密码

## 开发环境

为希望贡献代码或搭建本地开发环境的开发者提供：

### 🚀 快速开发环境搭建
```bash
# 克隆代码库
git clone https://github.com/veops/oneterm.git
cd oneterm/deploy

# 前端开发（实时编辑）
./dev-start.sh frontend

# 后端开发（实时编辑）  
./dev-start.sh backend

```

### 📖 详细开发指南
完整的环境搭建说明、故障排除和开发流程：
- **[开发环境搭建指南](deploy/DEV_README.zh.md)**

**环境要求**: Docker, Node.js 14.17.6+, Go 1.21.3+

## 代码贡献
我们欢迎所有开发者贡献代码，改善和扩展这个项目。请先阅读我们的[贡献指南](CONTRIBUTING_cn.md)。此外，您还可以通过社交媒体、活动和分享来支持 Veops 的开源。

<a href="https://github.com/veops/oneterm/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=veops/oneterm" />
</a>

## 更多开源
- [CMDB](https://github.com/veops/cmdb): 简单、轻量、通用的运维配置管理数据库
- [ACL](https://github.com/veops/acl): 一个简单通用的权限管理系统设计与实践。
- [messenger](https://github.com/veops/messenger): 一个简单轻量的消息发送服务。

## 与我联系
+ 邮箱: <a href="mailto:bd@veops.cn">bd@veops.cn</a>
+ 公众号：**维易科技OneOps**。关注后可以加入微信群，参与产品和技术交流   
  <img src="docs/images/wechat.png" alt="公众号: 维易科技OneOps" />