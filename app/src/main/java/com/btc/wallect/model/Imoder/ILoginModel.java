package com.btc.wallect.model.Imoder;



public interface ILoginModel {

    interface OnLoginListener {
        void onLoginSuccess();

        void onLoginFail();
    }

    void loginSubmit(String username, String password, OnLoginListener listener);
}
