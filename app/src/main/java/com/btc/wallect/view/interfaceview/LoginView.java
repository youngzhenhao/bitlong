package com.btc.wallect.view.interfaceview;

/**
 * UI接口
 */

public interface LoginView {
    /**
     * 显示进度条
     */
    void showProgress();

    /**
     * 隐藏进度条
     */
    void hideProgress();

    /**
     * 登录成功处理UI
     */
    void loginSuccess();

    /**
     * 登录失败处理UI
     */
    void loginFail();
}
