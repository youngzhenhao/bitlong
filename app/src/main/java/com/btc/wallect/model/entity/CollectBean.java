package com.btc.wallect.model.entity;

public class CollectBean {
    private String name;

    private int imageId;

    public CollectBean(String name, int imageId) {
        this.name = name;
        this.imageId = imageId;
    }

    public String getName() {
        return name;
    }

    public int getImageId() {
        return imageId;
    }

}
