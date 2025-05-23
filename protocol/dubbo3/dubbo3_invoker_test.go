/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dubbo3

const (
	mockDubbo3CommonUrl2 = "tri://127.0.0.1:20003/DubboGreeterImpl?accesslog=&anyhost=true&app.version=0.0.1&application=BDTService&async=false&bean.name=greeterImpl" +
		"&category=providers&cluster=failover&dubbo=dubbo-provider-golang-2.6.0&environment=dev&execute.limit=&execute.limit.rejected.handler=&generic=false&group=&interface=org.apache.dubbo.DubboGreeterImpl" +
		"&ip=192.168.1.106&loadbalance=random&methods.SayHello.loadbalance=random&methods.SayHello.retries=1&methods.SayHello.tps.limit.interval=&methods.SayHello.tps.limit.rate=&methods.SayHello.tps.limit.strategy=" +
		"&methods.SayHello.weight=0&module=dubbogo+say-hello+client&name=BDTService&organization=ikurento.com&owner=ZX&pid=49427&reference.filter=cshutdown&registry.role=3&remote.timestamp=1576923717&retries=" +
		"&service.filter=echo%2Ctoken%2Caccesslog%2Ctps%2Cexecute%2Cpshutdown&side=provider&timestamp=1576923740&tps.limit.interval=&tps.limit.rate=&tps.limit.rejected.handler=&tps.limit.strategy=&tps.limiter=&version=&warmup=100!"
)

//func TestInvoke(t *testing.T) {
//	go internal.InitDubboServer()
//	time.Sleep(time.Second * 3)
//
//	url, err := common.NewURL(mockDubbo3CommonUrl2)
//	assert.Nil(t, err)
//
//	invoker, err := NewDubboInvoker(url)
//
//	assert.Nil(t, err)
//
//	args := []reflect.Value{}
//	args = append(args, reflect.ValueOf(&internal.HelloRequest{Name: "request name"}))
//	bizReply := &internal.HelloReply{}
//	invo := invocation.NewRPCInvocationWithOptions(invocation.WithMethodName("SayHello"),
//		invocation.WithParameterValues(args), invocation.WithReply(bizReply))
//	res := invoker.Invoke(context.Background(), invo)
//	assert.Nil(t, res.Error())
//	assert.NotNil(t, res.Result())
//	assert.Equal(t, "Hello request name", bizReply.Message)
//}
//
//func TestInvokeTimoutConfig(t *testing.T) {
//	go internal.InitDubboServer()
//	time.Sleep(time.Second * 3)
//
//	// test for millisecond
//	tmpMockUrl := mockDubbo3CommonUrl2 + "&timeout=300ms"
//	url, err := common.NewURL(tmpMockUrl)
//	assert.Nil(t, err)
//
//	invoker, err := NewDubboInvoker(url)
//	assert.Nil(t, err)
//
//	assert.Equal(t, invoker.timeout, time.Duration(time.Millisecond*300))
//
//	// test for second
//	tmpMockUrl = mockDubbo3CommonUrl2 + "&timeout=1s"
//	url, err = common.NewURL(tmpMockUrl)
//	assert.Nil(t, err)
//
//	invoker, err = NewDubboInvoker(url)
//	assert.Nil(t, err)
//	assert.Equal(t, invoker.timeout, time.Duration(time.Second))
//
//	// test for timeout default config
//	url, err = common.NewURL(mockDubbo3CommonUrl2)
//	assert.Nil(t, err)
//
//	invoker, err = NewDubboInvoker(url)
//	assert.Nil(t, err)
//	assert.Equal(t, invoker.timeout, time.Duration(time.Second*3))
//}
