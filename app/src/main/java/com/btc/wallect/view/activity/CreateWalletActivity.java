package com.btc.wallect.view.activity;

import android.content.DialogInterface;
import android.content.Intent;
import android.os.Bundle;
import android.text.TextUtils;
import android.view.View;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import com.btc.wallect.presenter.IPresenter.ILoginPresenter;
import com.btc.wallect.presenter.compl.LoginPresentImpl;
import com.btc.wallect.R;
import com.btc.wallect.model.entity.UserInfo;
import com.btc.wallect.model.entity.Wallet;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.DialogUtil;
import com.btc.wallect.utils.UiUtils;
import com.btc.wallect.view.activity.base.BaseActivity;
import com.btc.wallect.view.interfaceview.LoginView;

import java.util.List;

import butterknife.BindView;
import butterknife.OnClick;


public class CreateWalletActivity extends BaseActivity implements LoginView {
    @BindView(R.id.ed_wallet_name)
    EditText name;
    @BindView(R.id.ed_input_passWord)
    EditText pwd;
    @BindView(R.id.ed_reuse_input_passWord)
    EditText reusePwd;
    @BindView(R.id.ed_remind_txt)
    EditText remindTxt;
    @BindView(R.id.img_passWord)
    ImageView imgPasswordSate;
    @BindView(R.id.tv_submit)
    TextView mTvSubmit;
//    @BindView(R.id.ed_remind_txt)
//    EditText edRemind;

    ILoginPresenter presenter;
    private boolean passWordState = true;
    private String page = "";

    @Override
    protected int setContentView() {
        return R.layout.activity_create_wallet;
    }

    @Override
    protected void init(View view, Bundle savedInstanceState) {
        setImgBack(true);
        Intent in = getIntent();
        page = in.getStringExtra(ConStantUtil.KEY_TOACTION);
        isCreateWallet();


        presenter = new LoginPresentImpl(this, new UserInfo());
        setImgPasswordSate();

    }

    @OnClick({R.id.tv_submit, R.id.img_passWord})
    public void submitClick(View view) {
        if (view.getId() == R.id.tv_submit) {
            String account = name.getText().toString().trim();
            if (TextUtils.isEmpty(account)) {
                DialogUtil.showSimpleDialog(this, "提示", "钱包名称不能为空", null);
                return;
            }
            if (account.length() > 12) {
                DialogUtil.showSimpleDialog(this, "提示", "钱包名称不能大于12个字", null);
                return;
            }

            String password = pwd.getText().toString().trim();
            String reuse_Pwd = reusePwd.getText().toString().trim();
            if (TextUtils.isEmpty(password)) {
                DialogUtil.showSimpleDialog(this, "提示", "请输入密码", null);
                return;
            }
            if (TextUtils.isEmpty(reuse_Pwd)) {
                DialogUtil.showSimpleDialog(this, "提示", "请输入重复密码", null);
                return;
            }
            if (!password.equals(reuse_Pwd)) {
                DialogUtil.showSimpleDialog(this, "提示", "密码不一致", null);
                return;
            }
            //    String txt = edRemind.getText().toString().trim();
            String txt = "111";
//            if (!TextUtils.isEmpty(account) && !TextUtils.isEmpty(password)){
//                presenter.loginSubmit(account, password);
//            }else {
//                DialogUtil.showSimpleDialog(this,"错误提示","账户密码不能为空",null);
//            }


            setDataSave(account, password, txt, "");
            if (isCreateWallet()){
                openActivity(CollectActivity.class);
            }else {
                openActivity(EditMnemonWordActivity.class);
            }


        } else if (view.getId() == R.id.img_passWord) {
            setImgPasswordSate();
        }
    }

    @Override
    public void showProgress() {
        DialogUtil.showProgress(this, "登陆中");
    }

    @Override
    public void hideProgress() {
        DialogUtil.dismissProgress();
    }

    @Override
    public void loginSuccess() {
        DialogUtil.showSimpleDialog(this, "登陆成功", "验证通过，是否进入到首页？", new DialogInterface.OnClickListener() {
            @Override
            public void onClick(DialogInterface dialogInterface, int i) {
                openActivity(MainActivity.class);
            }
        });
    }

    @Override
    public void loginFail() {
        DialogUtil.showSimpleDialog(this, "错误提示", "账户密码不正确", null);
    }

    private void setImgPasswordSate() {
        if (passWordState) {
            UiUtils.setHidePassword(CreateWalletActivity.this, reusePwd, imgPasswordSate, passWordState);
            passWordState = false;
        } else {
            UiUtils.setHidePassword(CreateWalletActivity.this, reusePwd, imgPasswordSate, passWordState);
            passWordState = true;
        }
    }

    /**
     * 保存数据
     */
    public void setDataSave(String name, String passWord, String txt, String collect) {
        try {

            Wallet wallet = new Wallet();
            wallet.name = name;
            wallet.password = passWord;
            wallet.txt = txt;
            wallet.collect = "";

            Long result = wallectDao.save(wallet);

            if (result != -1) {
                Toast.makeText(getApplication(), "保存数据成功!返回插入行号是[" + result + "]", Toast.LENGTH_SHORT).show();
                //  showToast("保存数据成功!返回插入行号是["+result+"]");
            } else {
                // showToast("保存数据失败result["+result+"]");
                Toast.makeText(getApplication(), "保存数据失败result[\"+result+\"]", Toast.LENGTH_SHORT).show();
            }

        } catch (Exception e) {
            e.printStackTrace();
        }
    }



    private boolean isCreateWallet() {

        if (page.equals(ConStantUtil.V_TOACTION_CREATE)) {
           // setTitle(R.string.app_txt_create_wallect);
            setTitle("创建钱包");
           // mTvSubmit.setText(R.string.app_btn_sure);
            mTvSubmit.setText("确认");
            return true;
        } else {
           // setTitle(R.string.app_txt_input_wallect);
            setTitle("导入钱包");
         //   mTvSubmit.setText(R.string.app_btn_sure_input);
            mTvSubmit.setText("确认导入");
            return false;
        }

    }


}
