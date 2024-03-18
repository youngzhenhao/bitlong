package com.btc.wallect.view.activity;

import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;

import com.btc.wallect.R;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.view.activity.base.BaseActivity;

import butterknife.BindView;
import butterknife.OnClick;

public class SelCreateWalletAct extends BaseActivity {


    @BindView(R.id.img_create_wallet)
    ImageView imgCreate;

    @Override
    protected int setContentView() {
        return R.layout.act_create_wallet;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setTitle(" ");
        setTitleHide(false);


    }

    @OnClick({R.id.img_create_wallet, R.id.img_input_wallet, R.id.img_hardWard_wallet, R.id.img_observer_wallet})
    public void onClick(View view) {
        if (view.getId() == R.id.img_create_wallet) {
            openActivityData(CreateWalletActivity.class, ConStantUtil.V_TOACTION_CREATE);
        } else if (view.getId() == R.id.img_input_wallet) {
            openActivityData(CreateWalletActivity.class, ConStantUtil.V_TOACTION_INPUT);
        } else if (view.getId() == R.id.img_hardWard_wallet) {

        } else if (view.getId() == R.id.img_observer_wallet) {
            openActivity(ImportKeyAcivity.class);
        }

    }


    private void toAction() {
        openActivity(CreateWalletActivity.class);

    }

    @Override
    protected void onDestroy() {

        super.onDestroy();
    }
}
