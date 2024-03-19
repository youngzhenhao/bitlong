package com.btc.wallect.view.activity;

import android.os.Bundle;
import android.view.View;

import com.btc.wallect.R;
import com.btc.wallect.view.activity.base.BaseActivity;

import java.util.Timer;
import java.util.TimerTask;

import io.reactivex.disposables.CompositeDisposable;

public class WelcomeAct  extends BaseActivity {


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
        openActivity(SelCreateWalletAct.class);
      //+  openActivity(MainActivity.class);
        finish();
    }

    @Override
    protected void onDestroy() {
        if (mCompositeDisposable != null) {
            mCompositeDisposable.dispose();
        }
        super.onDestroy();
    }
}
