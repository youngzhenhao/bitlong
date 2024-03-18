package com.btc.wallect.model.entity;

import android.os.Parcel;
import android.os.Parcelable;

public class TaprTabListBean {
   private String dealName;
    private double dealAmount;
    private  double dealPrice;
    private double dealdDetailPrice;
    private  double rose;
    private int state;

    public String getDealName() {
        return dealName;
    }

    public void setDealName(String dealName) {
        this.dealName = dealName;
    }

    public double getDealAmount() {
        return dealAmount;
    }

    public void setDealAmount(double dealAmount) {
        this.dealAmount = dealAmount;
    }

    public double getDealPrice() {
        return dealPrice;
    }

    public void setDealPrice(double dealPrice) {
        this.dealPrice = dealPrice;
    }

    public double getDealdDetailPrice() {
        return dealdDetailPrice;
    }

    public void setDealdDetailPrice(double dealdDetailPrice) {
        this.dealdDetailPrice = dealdDetailPrice;
    }

    public double getRose() {
        return rose;
    }

    public void setRose(double rose) {
        this.rose = rose;
    }

    public int getState() {
        return state;
    }

    public void setState(int state) {
        this.state = state;
    }
}



