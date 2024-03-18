package com.btc.wallect.model.entity;

import android.os.Handler;

import com.btc.wallect.model.Imoder.ILoginModel;


public class UserInfo implements ILoginModel {
    private String age;
    private String name;
    private String gender;
    private String hobby;

    public String getAge() {
        return age;
    }

    public void setAge(String age) {
        this.age = age;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getGender() {
        return gender;
    }

    public void setGender(String gender) {
        this.gender = gender;
    }

    public String getHobby() {
        return hobby;
    }

    public void setHobby(String hobby) {
        this.hobby = hobby;
    }


    @Override
    public void loginSubmit(final String username, final String password, final OnLoginListener listener) {

        //此次Handler制造延迟1秒和回调——纯属模仿网络请求。实际开发可以替换
        new Handler().postDelayed(new Runnable() {
            @Override
            public void run() {
                if (username.equals("a") && password.equals("a")) {
                    listener.onLoginSuccess();
                } else {
                    listener.onLoginFail();
                }
            }
        }, 1000);
    }
}
