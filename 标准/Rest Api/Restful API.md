# Restful API

某天听见leader和其他部门人员的对话，老大说你们可以使用restful访问我们的接口，那个哥们竟然说他不知道restful api是什么。老大当时气的冒烟，你这一个程序员不知道restful是什么，那我们今天这个对话没法继续谈了，没有继续下去的意义。

那个时候想了一下，我虽然每天都在使用restful风格，但是要是真的具体说它是什么，我还真的说不出来，所以今天就学习了一下。

### 正文

 REST就是Representational State Transfer  翻译过来就是表诉状态转移，只是一种规定的风格，只要符合这种风格的都可以叫做restful风格

借鉴这个故事(别人的)来描述一下restful的几个级别

## 故事

老婆经常喜欢翻看我订阅的技术杂志，她总能从她的视角提出很多有趣的问题。

一个悠闲的周日下午，她午觉醒来，又习惯性的抓起这个月的杂志，饶有兴趣地看了起来。

果不其然，看着看着，她又对我发难了，“Restful是什么呀，老公？是restaurant的形容词吗，突然就觉得好饿了啊......”

作为一个合格的程序员，我一直把能够将一项技术讲给老婆听，并且能给她讲懂，作为我已经掌握了这项技术的标准。

如果我直接回答说，“REST就是Representational State Transfer的缩写呀，翻译为中文就是‘表述性状态转移’”，那她今晚肯定得罚我跪键盘。我必须找个合适的机会，把Restful的**来龙去脉**给她形象的描述一遍。

“走，咱们去楼下咖啡厅吃个下午茶吧”，我对老婆说。

“一个芝士蛋糕，一杯拿铁，两条吸管，谢谢”，我对前台的服务员说，然后我们找了个角落坐了下来。

## Level 0 - 面向前台

“刚才我们向前台点了一杯拿铁，这个过程可以用这段文字来描述”，说着，我在纸上写下了这段JSON，虽然她不知道什么叫JSON，但理解这段文字对于英语专业8级的她，实在再简单不过。

```text
{
    "addOrder": {
        "orderName": "latte"
    }
}
```

“我们通过这段文字，告诉前台，新增一笔订单，订单是一杯拿铁咖啡”，接着，前台给我们返回这么一串回复：

```text
{
    "orderId": "123456"
}
```

“订单ID？还是订单编号？”

“恩恩，就是订单编号”

“那我们就等着前台喊‘订单123456的客户可以取餐了’，然后就可以开吃了！”

“哈哈，你真聪明，不过，在这之前，假设我们有一张会员卡，我们想查询一下这张会员卡的余额，这时候，要向前台发起另一个询问”，我继续在纸上写着：

```text
{
    "queryBalance": {
        "cardId": "886333"
    }
}
```

“查询卡号为886333的卡的余额？”

“真棒！接着，查询的结果返回来了”

```text
{
    "balance": "0"
}
```

“切，没钱......”

“哈哈，没钱，现在我们要跟前台说，这杯咖啡不要了”，我在纸上写到：

```text
{
    "deleteOrder": {
        "orderId": "123456"
    }
}
```

“哼，这就把订单取消啦？”

## Level 1 - 面向资源

“现在这家咖啡店越做越大，来喝咖啡的人越来越多，单靠前台显然是不行的，店主决定进行分工，每个**资源**都有专人负责，我们可以直接**面向资源**操作。”

"面向资源？”

“是的，比如还是下单，请求的内容不变，但是我们多了一条消息”，我在纸上画出这次的模型：

```text
/orders

{
    "addOrder": {
        "orderName": "latte"
    }
}
```

“多了一个斜杠和orders？这是什么意思？”

“这个表示我们这个请求是发给哪个资源的，订单是一种资源，我们可以理解为是咖啡厅专门管理订单的人，他可以帮我们处理所有有关订单的操作，包括新增订单、修改订单、取消订单等操作”

“Soga...”

“接着还是会返回订单的编号给我们”

```text
{
    "orderId": "123456"
}
```

“下面，我们还是要查询会员卡余额，这次请求的资源变成了cards”

```text
/cards

{
    "queryBalance": {
        "cardId": "886333"
    }
}
```

“接下来是取消订单”

“这个我会”，说着，她抢走我手上的笔，在纸上写了起来：

```text
/orders

{
    "deleteOrder": {
        "orderId": "123456"
    }
}
```

## Level 2 - 打上标签

“接下来，店主还想继续优化他的咖啡厅的服务流程，他发现负责处理订单的员工，每次都要去订单内容里面看是新增订单还是删除订单，还是其他的什么操作，十分不方便，于是规定，所有新增资源的请求，都在请求上面写上大大的‘POST’，表示这是一笔新增资源的请求”

“其他种类的请求，比如查询类的，用‘GET’表示，删除类的，用‘DELETE’表示”

“还有修改类的，修改分为两种，第一种，如果这个修改，**无论发送多少次，最后一次修改后的资源，总是和第一次修改后的一样**，比如将拿铁改为猫屎，那么用‘PUT’表示；第二种，如果这个修改，**每次修改都会让这个资源和前一次的不一样**，比如是加一杯咖啡，那么这种请求用‘PATCH’或者‘POST’表示”，一口气讲了这么多，发现她有点似懂非懂。

“来，我们再来重复上面那个过程，来一杯拿铁”，我边说边画着：

```text
POST /orders

{
    "orderName": "latte"
}
```

"请求的内容简洁多啦，不用告诉店员是addOrder，看到POST就知道是新增"，她听的很认真，理解的也很透彻。

"恩恩，返回的内容还是一样"

```text
{
    "orderId": "123456"
}
```

“接着是查询会员卡余额，这次也简化了很多”

```text
GET /cards

{
    "cardId": "886333"
}
```

“这个请求我们还可以进一步优化为这样”

```text
GET /cards/886333
```

“Soga，直接把要查询的卡号写在后面了”

“没错，接着，取消订单”

```text
DELETE /orders/123456
```

## Level 3 - 完美服务

“忽然有一天，有个顾客抱怨说，他买了咖啡后，不知道要怎么取消订单，咖啡厅一个店员回了一句，你不会看我们的宣传单吗，上面不写着：

```text
DELETE /orders/{orderId}
```

顾客反问道，谁会去看那个啊，店员不服，又说到，你瞎了啊你......据说后面两人吵着吵着还打了起来...”

“噗，真是悲剧...”

“有了这次教训，店长决定，顾客下了单之后，不仅给他们返回订单的编号，还给顾客返回所有可以对这个订单做的操作，比如告诉用户如何删除订单。现在，我们还是发出请求，请求内容和上一次一样”

```text
POST /orders

{
    "orderName": "latte"
}
```

“但是这次返回时多了些内容”

```text
{
    "orderId": "123456",
    "link": {
        "rel": "cancel",
        "url": "/order/123456"
    }
}
```

“这次返回时多了一项link信息，里面包含了一个rel属性和url属性，rel是relationship的意思，这里的关系是cancel，url则告诉你如何执行这个cancel操作，接着你就可以这样子来取消订单啦”

```text
DELETE /orders/123456
```

“哈哈，这服务真是贴心，以后再也不用担心店员和顾客打起来了”

“订单123456的客户可以取餐了”，伴随着咖啡厅的广播，我们吃起了下午茶，一杯拿铁，两支吸管...... 

## 总结

上面这个故事真的太清晰了，但是还需要总结一下，restful风格就是让每一条请求带着一个请求的状态，如，GET，POST，PUT，DELETE，还要带着要访问的资源，联系上文就是order/ 或者 cards/，最后再带着自己的内容，这里没有说到lever3，因为一般的web，api就只到level2。

## 附加

关于restful有一个幂等性的概述，即GET、PUT、DELETE是幂等的，而POST、PATCH不是幂等的。

HTTP幂等方法，是指无论调用多少次都不会有不同结果的 HTTP 方法。不管你调用一次，还是调用一百次，一千次，结果都是相同的。 值得注意，幂等性指的是作用于结果而非资源本身。怎么理解呢？例如，这个HTTP GET方法可能会每次得到不同的返回内容，但并不影响资源。 