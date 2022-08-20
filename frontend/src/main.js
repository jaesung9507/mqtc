import './style.css';

import {MqttConnect, MqttSubscription, MqttPublish} from '../wailsjs/go/main/App';
import {EventsOn} from '../wailsjs/runtime/runtime';

window.OnConnectButton = function() {
    try {
        MqttConnect("")
            .then((result) => {
                if(result) {
                    const connectStatus = document.getElementById("connectStatus");
                    connectStatus.style.backgroundColor = "green";
                }
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

window.OnPublishButton = function() {
    const topic = document.getElementById("pubTopic").value;
    const msg = document.getElementById("pubMsg").value;
    try {
        MqttPublish(topic, msg)
            .then((result) => {
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
}

window.OnSubscriptionButton = function() {
    const topic = document.getElementById("subTopic").value;
    if(topic == "") {
        return;
    }
    try {
        MqttSubscription(topic)
            .then((result) => {
                if(result) {
                    const subList = document.getElementById("subList");
                    let item = document.createElement("option");
                    item.text = topic;
                    subList.options.add(item);
                }
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

EventsOn("OnMqttMessage", function(topic, payload, qos, retained) {
    const row = document.createElement("div");
    row.className = "subMsgItem";

    const topicDiv = document.createElement("div");
    topicDiv.className = "subMsgTopic";
    topicDiv.innerText = topic;
    row.appendChild(topicDiv);

    const timeDiv = document.createElement("div");
    timeDiv.className = "subMsgTime";
    timeDiv.innerText = "YYYY MM DD";
    row.appendChild(timeDiv);

    const payloadDiv = document.createElement("div");
    payloadDiv.className = "subMsgPayload";
    payloadDiv.innerText = payload;
    row.appendChild(payloadDiv);

    const qosDiv = document.createElement("div");
    qosDiv.className = "subMsgQos";
    qosDiv.innerText = qos;
    row.appendChild(qosDiv);

    const retainedDiv = document.createElement("div");
    retainedDiv.className = "subMsgRetain";
    retainedDiv.innerHTML = retained;
    row.appendChild(retainedDiv);

    const subMsg = document.getElementById("subMsg");
    subMsg.appendChild(row);
});

EventsOn("OnMqttDisconnect", function() {
    const connectStatus = document.getElementById("connectStatus");
    connectStatus.style.backgroundColor = "red";
})
