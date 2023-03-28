package com.company;

import org.eclipse.paho.client.mqttv3.*;

import java.nio.charset.StandardCharsets;
import java.util.Arrays;

public class Main {

    public static void main(String[] args) throws MqttException {
        String broker       = "tcp://test.mosquitto.org:1883";
        MqttClient Client = new MqttClient(broker, MqttClient.generateClientId());
        Client.setCallback(new MqttCallback() {
            @Override
            public void connectionLost(Throwable throwable) {
                System.out.println("Connection to MQTT broker lost!");
            }

            @Override
            public void messageArrived(String s, MqttMessage mqttMessage) throws Exception {
                System.out.println(mqttMessage);
                String tmp = new String(mqttMessage.getPayload());
                double[] numArr = Arrays.stream(tmp.split(" ")).mapToDouble(Double::parseDouble).toArray();
                double div;
                div = numArr[0] * numArr[0] + numArr[1] * numArr[1] + numArr[2] * numArr[2];
                String sent = "";
                for(int i = 0; i < 3; i++) {
                    numArr[i] = numArr[i] / Math.sqrt(div);
                    sent += String.valueOf(numArr[i]) + " ";
                }
                System.out.println(div);
            }

            @Override
            public void deliveryComplete(IMqttDeliveryToken iMqttDeliveryToken) {

            }
        });
        Client.connect();
        Client.subscribe("IU/9");
    }
}
