/*
 * Distrubuted AI
 *
 * 分布式AI系统架构与API文档说明  系统架构图在https://github.com/lzhou1110/DistributedAI/blob/master/%E7%B3%BB%E7%BB%9F%E6%9E%B6%E6%9E%84.jpg  1.存储采用IPFS，可以使用storage/post(file_path)将一个文件上传至IPFS，并且返回文件在IPFS上的地址 使用/storage/{IPFSAddress}可以从地址`IPFSAddress`上下载文件  2.交易前： ①数据方： 数据方上传前在提供 （1）DataSchema：数据格式 （2）MetaSchema：原数据，主要是QI，用于客观描述数据 （3）Payment：数据售价 然后上传IPFS网络，用tansaction/bidData将数据====>数据池 参数和响应值可以看下方的API，如参数为\"DataSchemaAddress\"， 可以查看Model中的DataSchema方便理解。  ②运算方：运算方在上传前提供 （1）OperationSchemaAddress：运算类型地址 （2）ComputingAddress：运算资源地址 （3）PaymentAddress：运算价格地址 （4）ComputerAttributesAddress：运算资源信息 然后调用tansaction/bidComputing====> 运算池 并将运算方公钥上传IPFS，并将地址上传。  ③ 模型方：公开同态加密的公钥，用公钥加密模型，上传IPFS网络，并将两者地址上传  3.开始交易 （1）模型方浏览当前的数据，自行挑选数据（提供两种方式：1 输入schema 地址 -> 匹配所有数据地址 2 输入所需数据的标签 -> 匹配schema 地址），将符合的数据的Meta Data下载至本地，利用Meta Data中的数据计算数据质量，选定数据 （2）调用API：/transaction/askTraining 发起交易，返回TransationDetail（json,目前只有ID一个属性） （3）发起交易成功则调用/transaction/askForData，向数据方请求上传数据。（具体参数看API和Model） （4）数据方先用模型方的公钥进行同态加密，再用运算方的AES公钥加密，上传至IPFS，并将地址上传 （5）运算方下载数据，进行AES解密，得到同态加密的模型和数据，开始训练（运算方应该有一个代码托管的框架，直接调用接口train便可以训练任何模型）。 （6）达到stopCondition停止训练，并调用/transaction/uploadTrainResult将运算结果上传，通过协议层发送至模型方，模型方解密收到运算结果并进行整合。  4.结束交易。 （1）模型方根据交易时所确定的Token Strategy（如:反向博弈论）进行tokens的分配 （2）模型的使用：模型的使用必须针对加密过的数据，也就是使用模型时需要先用同态加密对数据进行加密，然后才能使用模型。
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"
	"go-ethereum/accounts/abi/bind"
	"go-ethereum/accounts/abi/bind/backends"
	"go-ethereum/core"
	"go-ethereum/crypto"


	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "./go"
)

func main() {
	log.Printf("Server started")




	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
