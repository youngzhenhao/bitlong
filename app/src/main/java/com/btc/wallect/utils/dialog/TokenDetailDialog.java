package com.btc.wallect.utils.dialog;

import android.app.Dialog;
import android.content.Context;
import android.os.Bundle;
import android.view.Display;
import android.view.Gravity;
import android.view.View;
import android.view.Window;
import android.view.WindowManager;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;

import com.btc.wallect.R;
import com.btc.wallect.model.entity.WalletListBean;

import java.util.List;

public class TokenDetailDialog extends Dialog {
    Context mContext;
    private ImageView mImg_out_query,mImg_in_query,mImg_close;
    private ImageView mImg_add2;
    private TextView mTv1, mTv2, mTv3, mTv4, mTv5, mTv6, mTv7, mTv8, mTv9;


    private TextView mTv_submit;


    public List<WalletListBean> walletList;

    public TokenDetailDialog(@NonNull Context context) {
        super(context, R.style.wallectDialog);
        this.mContext=context;

    }








    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.dia_token_detail);
        //空白处不能取消动画
        setCanceledOnTouchOutside(false);

        //初始化界面控件
        initView();

        //初始化界面数据
        initData();

    }

    /**
     * 初始化界面控件
     */
    private void initView() {
        mTv1 = findViewById(R.id.tv_1);
        mTv2 = findViewById(R.id.tv_tv2);
        mTv3 = findViewById(R.id.tv_3);
        mTv4 = findViewById(R.id.tv_4);
        mTv5 = findViewById(R.id.tv_5);
        mTv6 = findViewById(R.id.tv_6);
        mTv7 = findViewById(R.id.tv_7);
        mTv8 = findViewById(R.id.tv_8);
        mTv9 = findViewById(R.id.tv_9);

        mImg_close=findViewById(R.id.img_close);



        Window window = getWindow();
        window.setLayout(WindowManager.LayoutParams.MATCH_PARENT, WindowManager.LayoutParams.WRAP_CONTENT);
        window.setGravity(Gravity.BOTTOM);
        WindowManager m = getWindow().getWindowManager();
        Display d = m.getDefaultDisplay();
        WindowManager.LayoutParams p = getWindow().getAttributes();
        p.width = d.getWidth(); //设置dialog的宽度为当前手机屏幕的宽度
        getWindow().setAttributes(p);

    }

    /**
     * 初始化界面控件的显示数据
     */
    private void initData() {





        mImg_close.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                dismiss();

            }
        });
    }



    public interface onOutClickListener {
        public void onOutClick();
    }
    public interface onInClickListener {
        public void onInClick();
    }
    public interface onSureClickListener {
        public void onOkClick();
    }
}
