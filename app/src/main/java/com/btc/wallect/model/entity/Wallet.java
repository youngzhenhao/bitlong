package com.btc.wallect.model.entity;

import java.io.Serializable;

public class Wallet implements Serializable {
    public Long id;
    public String name;
    public String password;
    public String txt;
    public String collect;

    public String creatDate;

    //头像
    public byte[] logoHead;
    @Override
    public String toString() {
        return "Student{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", sex='" + password + '\'' +
                ", age=" + txt +
                ", clazz='" + collect + '\'' +
                ", creatDate='" + creatDate + '\'' +
                '}';
    }


}