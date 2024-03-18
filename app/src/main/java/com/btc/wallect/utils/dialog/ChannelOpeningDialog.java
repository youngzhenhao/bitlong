package com.btc.wallect.utils.dialog;

import android.app.Dialog;
import android.content.Context;
import android.os.Bundle;
import android.view.Display;
import android.view.Gravity;
import android.view.View;
import android.view.Window;
import android.view.WindowManager;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;

import com.btc.wallect.model.entity.WalletListBean;
import com.btc.wallect.R;

import java.util.List;

public class ChannelOpeningDialog extends Dialog {
    Context mContext;
    private ImageView mImg_out_query,mImg_in_query,mImg_close;
    private ImageView mImg_add2;
    private EditText mEdOutbount,mEdInbount;
    private TextView mTvSetvice,mTvAbsenteeism;
    private TextView mTv_submit;

    private onOutClickListener OutClickListener;
    private onInClickListener InClickListener;
    private onSureClickListener onSureClickListener;

    public List<WalletListBean> walletList;

    public ChannelOpeningDialog(@NonNull Context context) {
        super(context, R.style.wallectDialog);
        this.mContext=context;

    }



    public void setAddOnclickListener( onOutClickListener outClick) {
        this.OutClickListener = outClick;
    }
    public void setAddOnclickListener( onInClickListener inClick) {
        this.InClickListener = inClick;
    }

    public void setOnSureClickListenerr( onSureClickListener inClick) {
        this.onSureClickListener = inClick;
    }





    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.dia_openning_channel);
        //空白处不能取消动画
        setCanceledOnTouchOutside(false);

        //初始化界面控件
        initView();

        //初始化界面数据
        initData();
        //初始化界面控件的事件
        initEvent();
    }

    /**
     * 初始化界面控件
     */
    private void initView() {
        mEdOutbount=findViewById(R.id.ed_outbount);
        mEdInbount=findViewById(R.id.ed_in_bount);
        mTvSetvice=findViewById(R.id.tv_setvice);
        mTvAbsenteeism=findViewById(R.id.tv_absenteeism);
        mTv_submit=findViewById(R.id.tv_submit);
        mImg_close=findViewById(R.id.img_close);
        mImg_out_query=findViewById(R.id.img_out_query);
        mImg_in_query=findViewById(R.id.img_in_query);


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



    }

    /**
     * 初始化界面的确定和取消监听
     */
    private void initEvent() {
        mImg_out_query.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (OutClickListener != null) {
                    OutClickListener.onOutClick();
                }
            }
        });

        //设置取消按钮被点击后，向外界提供监听
        mImg_in_query.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (InClickListener != null) {
                    InClickListener.onInClick();
                }
            }
        });
        mTv_submit.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (onSureClickListener != null) {
                    onSureClickListener.onOkClick();
                }
            }
        });
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
