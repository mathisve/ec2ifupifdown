package ec2ifupifdown

// hello rico!

// hello mathis!

// how are you doing on this fine day

// I am doing well, I had a single lesson today
// how are you doing? how is america?

// that is good to hear, I am doing splendid herre in the USA.

// what does this program do?

// controls ifupifdown for ec2 instances

// is that to do with resource scaling? like checking if a resource is idle or smth?

// no, allows you to shut down a network interface on an ec2 instance using golang
// wrote it for Berty a while back, it's pretty shit

// so it can be like a kill switch?

// yea

// neat
// some of the code is questionable haha

// it really is, was hacked together
// the internal API is really wackadoodle

// the one from aws? wdym by "internal api"?

// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html
// read this
// you can query this API to see your EC2 instances metadata, like network interfaces, mac adresses, security groups, etc.

// so you query this api and then use the result

// pretty much
// im going to commit this conversation to master as "added conversational documentation"

// lol
// so is it literally just parsing the response from the api, and then allowing you
// to up/down the network interface for each result?

// exactly, couldn't figure out a better way to do it whithout using the "external" AWS api to see the instances metadata
// it's very hacky but it works "really well"

// if it works, it works

// famous last words
// can I add you as a co-author on github?

// sure, if you want haha

// yes sir
// watch this, i'm gonna start a revolution

// about to commit

// LGTM!

// ill make it a branch, so we can make a PR and review it, etc.


// where is this going to be PR'd against?

// what is your github name and email?

// getaddrinfo <oss@getaddrinfo.com>

// 👍