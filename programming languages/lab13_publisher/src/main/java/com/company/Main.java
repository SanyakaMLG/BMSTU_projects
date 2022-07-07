package com.company;

import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;

public class Main {

    public static void main(String[] args) throws MqttException {
        String broker = "tcp://test.mosquitto.org:1883";
        MqttClient client = new MqttClient(broker, MqttClient.generateClientId());
        client.connect();
        client.publish("IU/9", new MqttMessage("10 0 0".getBytes()));
        client.subscribe("IU/9");
//        while (true) {
//            client.publish("MQTT Examples", new MqttMessage("Anatoly Zadvornyh donkey".getBytes()));
//        }
    }
}
