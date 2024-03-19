package com.btc.wallect.model.Imoder;

public interface IWallectFagmentModel {
    interface OnListener {
        void onLoginSuccess();

        void onLoginFail();
    }
    void loginSubmit( IWallectFagmentModel.OnListener listener);
}
