package com.btc.wallect.model.entity;

public class MainDateilBean {
    private  String btcName ;
    private String btcMode;
    private double btcAmount;
    private double btcAll;

    public String getBtcName() {
        return btcName;
    }

    public void setBtcName(String btcName) {
        this.btcName = btcName;
    }

    public String getBtcMode() {
        return btcMode;
    }

    public void setBtcMode(String btcMode) {
        this.btcMode = btcMode;
    }

    public double getBtcAmount() {
        return btcAmount;
    }

    public void setBtcAmount(double btcAmount) {
        this.btcAmount = btcAmount;
    }

    public double getBtcAll() {
        return btcAll;
    }

    public void setBtcAll(double btcAll) {
        this.btcAll = btcAll;
    }
}
