package com.btc.wallect.presenter.compl;

import com.btc.wallect.model.Imoder.ILoginModel;
import com.btc.wallect.model.Imoder.IWallectFagmentModel;
import com.btc.wallect.presenter.IPresenter.ILoginPresenter;
import com.btc.wallect.presenter.IPresenter.IWallectFragmentPresentImpl;
import com.btc.wallect.view.interfaceview.WallectFragmentView;

public class WallectFragmentPresentImpl implements IWallectFragmentPresentImpl, IWallectFagmentModel.OnListener{
    WallectFragmentView wallectFragmentView;
    @Override
    public void onLoginSuccess() {

    }

    @Override
    public void onLoginFail() {

    }

    @Override
    public void uidata() {
        wallectFragmentView.showUIData();
    }
}
