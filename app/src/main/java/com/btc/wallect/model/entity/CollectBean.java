package com.btc.wallect.model.entity;

public class CollectBean {
    private String name;

    private int imageId;
    public boolean isvarify;

    public CollectBean(String name, int imageId,boolean varify) {
        this.name = name;
        this.imageId = imageId;
        this.isvarify=varify;
    }

    public String getName() {
        return name;
    }

    public int getImageId() {
        return imageId;
    }

    public boolean isVarify() {
        return isvarify;
    }
}
