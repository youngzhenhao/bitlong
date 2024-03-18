package com.btc.wallect.view.activity.base;

import android.content.Intent;
import android.graphics.Color;
import android.os.Bundle;
import android.text.TextUtils;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.FrameLayout;
import android.widget.ImageView;
import android.widget.RelativeLayout;
import android.widget.TextView;

import com.btc.wallect.db.DBOpenHelper;
import com.btc.wallect.db.DBdao;
import com.btc.wallect.R;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.DialogUtil;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.trello.rxlifecycle2.LifecycleTransformer;
import com.trello.rxlifecycle2.components.support.RxAppCompatActivity;

import butterknife.ButterKnife;
import butterknife.Unbinder;


public abstract class BaseActivity<T extends BaseConstract.IBasePersenter> extends RxAppCompatActivity implements BaseConstract.IBaseView {
    FrameLayout mRootView;
    RelativeLayout rl;
    TextView mTitle;
    ImageView mImg_back;
    private Unbinder unbinder;

    protected T mPresenter;
    private DBOpenHelper dBhelpUtil;
    public DBdao studentDao;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_base);
        createView();
        init(mRootView, savedInstanceState);
        attachView();
        dBhelpUtil = new DBOpenHelper(this, DBOpenHelper.DB_NAME, null, DBOpenHelper.DB_VERSION);
        studentDao = new DBdao(this, dBhelpUtil);
        SharedPreferencesHelperUtil.getInstance().init(this);
    }

    private void createView() {
        rl = findViewById(R.id.title_rl);
        mTitle = findViewById(R.id.title);
        mImg_back = findViewById(R.id.img_back);
        mRootView = findViewById(R.id.view_content);
        View view = LayoutInflater.from(this).inflate(setContentView(), mRootView);
        unbinder = ButterKnife.bind(this, view);

    }

    private void attachView() {
        if (mPresenter != null) {
            mPresenter.attachView(this);
        }
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        unbinder.unbind();
        if (mPresenter != null) {
            mPresenter.detachView();
        }

    }

    /**
     * 设置布局
     *
     * @return
     */
    protected abstract int setContentView();

    /**
     * 子类初始化
     *
     * @param savedInstanceState
     */
    protected abstract void init(View view, Bundle savedInstanceState);

    /**
     * 设置标题
     *
     * @param str
     */
    protected void setTitle(String str) {
        rl.setVisibility(View.VISIBLE);
        if (str != null && !TextUtils.isEmpty(str)) {
            TextView title_view = findViewById(R.id.title);
            title_view.setText(str);
        }


    }

    protected void setTitleHide(boolean isShow) {
        if (isShow) {
            rl.setVisibility(View.VISIBLE);
        } else {
            rl.setVisibility(View.GONE);
        }


    }

    protected void setTitleBackgroundColor(int color) {
        if (color == 1) {
            rl.setBackgroundColor(Color.parseColor("#ffffffff"));
        } else if (color == 2) {

            rl.setBackgroundColor(Color.parseColor("#F5F5F5"));

        } else if (color == 3) {
            rl.setBackgroundColor(Color.parseColor("#665AF0"));

        } else if (color == 4) {
            rl.setBackgroundColor(Color.parseColor("#665AF0"));

        }

    }

    protected void setTitleTxtColor(int color) {
        if (color == 0) {
            mTitle.setTextColor(Color.parseColor("#FFFFFF"));
            mImg_back.setImageResource(R.mipmap.img_white_back);
        } else if (color == 2) {
            mTitle.setTextColor(Color.parseColor("#383838"));
            mImg_back.setImageResource(R.mipmap.img_back);
        }

    }

    /**
     * 显示返回按钮
     *
     * @param show
     */
    protected void setTextBack(boolean show) {
        TextView back_view = findViewById(R.id.back);
        if (show) {
            back_view.setVisibility(View.VISIBLE);
            back_view.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View view) {
                    finish();
                }
            });
        }

    }

    protected void setImgBack(boolean show) {
        ImageView back_view = findViewById(R.id.img_back);
        if (show) {
            back_view.setVisibility(View.VISIBLE);
            back_view.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View view) {
                    finish();
                }
            });
        }

    }

    /**
     * 普通页面跳转
     */
    protected void openActivity(Class<?> Action) {
        startActivity(new Intent(this, Action));
    }

    protected void openActivityData(Class<?> Action, String toWhere) {
        Intent intent = new Intent(this, Action);
        intent.putExtra(ConStantUtil.KEY_TOACTION, toWhere);

        startActivity(intent);
    }

    @Override
    public void showProgress() {
        DialogUtil.showProgress(this, "请求中。。");
    }

    @Override
    public void hideProgress() {
        DialogUtil.dismissProgress();
    }

    @Override
    public void showFaild(String onError) {
        DialogUtil.showSimpleDialog(this, "错误提示", onError, null);
    }

    //接口隔离原则：使用多个隔离的接口，比使用单个接口要好，目的就是降低类之间的耦合度，便于软件升级和维护。
    @Override
    public <T> LifecycleTransformer<T> bindToLife() {
        return this.<T>bindToLifecycle();
    }
}

