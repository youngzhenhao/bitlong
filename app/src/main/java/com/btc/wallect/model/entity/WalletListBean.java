package com.btc.wallect.model.entity;

public class WalletListBean  {


    public  String WallectName;
    private String WallectKey;
    private Double WallectAmount;
    private int Type;

    public String getWallectName() {
        return WallectName;
    }

    public void setWallectName(String wallectName) {
        WallectName = wallectName;
    }

    public String getWallectKey() {
        return WallectKey;
    }

    public void setWallectKey(String wallectKey) {
        WallectKey = wallectKey;
    }

    public Double getWallectAmount() {
        return WallectAmount;
    }

    public void setWallectAmount(Double wallectAmount) {
        WallectAmount = wallectAmount;
    }

    public int getType() {
        return Type;
    }

    public void setType(int type) {
        Type = type;
    }
}
