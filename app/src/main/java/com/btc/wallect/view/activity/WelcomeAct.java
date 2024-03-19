package com.btc.wallect.view.activity;

import android.os.Bundle;
import android.view.View;

import com.btc.wallect.R;
import com.btc.wallect.model.entity.Wallet;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.LogUntil;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.google.gson.Gson;

import java.util.List;
import java.util.Timer;
import java.util.TimerTask;

import io.reactivex.disposables.CompositeDisposable;

public class WelcomeAct extends BaseActivity {


    CompositeDisposable mCompositeDisposable = new CompositeDisposable();

    @Override
    protected int setContentView() {
        return R.layout.act_welcome;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {


        Timer timer = new Timer();
        TimerTask task = new TimerTask() {
            @Override
            public void run() {
                toMain();
            }
        };
        timer.schedule(task, 1500); //1.5秒后跳转

    }


    private void toMain() {
        if (mCompositeDisposable != null) {
            mCompositeDisposable.dispose();
        }
        getWallectList();

        //+  openActivity(MainActivity.class);
        finish();
    }

    private void getWallectList() {
        boolean isWallect = SharedPreferencesHelperUtil.getInstance().getBooleanValue(ConStantUtil.ISWALLECT, false);
        if (isWallect) {
            List<Wallet> walletList = selectWallectData();
            LogUntil.d(new Gson().toJson(walletList));
            for (int i = 0; i < walletList.size(); i++) {

            }
            openActivity(MainActivity.class);
        } else {
            openActivity(SelCreateWalletAct.class);
        }


    }

    @Override
    protected void onDestroy() {
        if (mCompositeDisposable != null) {
            mCompositeDisposable.dispose();
        }
        super.onDestroy();
    }
}
