package com.btc.wallect.presenter.compl;

import com.btc.wallect.model.Imoder.ILoginModel;
import com.btc.wallect.presenter.IPresenter.ILoginPresenter;
import com.btc.wallect.view.interfaceview.LoginView;

/**
 * 中介者——处理视图和模型
 */

public class LoginPresentImpl implements ILoginPresenter, ILoginModel.OnLoginListener {

    private LoginView mView;
    private ILoginModel mModel;

    /**
     * 构造函数进行实例化
     * @param mView
     * @param mModel
     */
    public LoginPresentImpl(LoginView mView, ILoginModel mModel) {
        this.mView = mView;
        this.mModel = mModel;
    }

    /**
     * 登陆方法
     * @param username
     * @param password
     */
    @Override
    public void loginSubmit(String username, String password) {
        mView.showProgress();
        mModel.loginSubmit(username,password,this);//在模型中通过接口分解结果，并将回调方法暴露出来，可以做到抽离开Model对View的联系。
    }

    @Override
    public void onLoginSuccess() {
        mView.hideProgress();
        mView.loginSuccess();
    }

    @Override
    public void onLoginFail() {
        mView.loginFail();
        mView.hideProgress();
    }
}
