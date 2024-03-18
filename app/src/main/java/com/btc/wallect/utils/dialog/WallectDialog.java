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
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.R;
import com.btc.wallect.adapter.WallectListAdapter;
import com.btc.wallect.model.entity.WalletListBean;

import java.util.List;

public class WallectDialog extends Dialog {
    Context mContext;
    private ImageView mImg_close;
    private ImageView mImg_add2;
    private RecyclerView mRecycler_wallect;
    private TextView titleTV;//消息标题文本
    private TextView message;//消息提示文本
    private String titleStr;//从外界设置的title文本
    private String messageStr;//从外界设置的消息文本
    //确定文本和取消文本的显示的内容
    private String yesStr, noStr;
    private onAddClickListener addClickListener;

    public List<WalletListBean> walletList;

    public WallectDialog(@NonNull Context context, List<WalletListBean> walletListBeans) {
        super(context, R.style.wallectDialog);
        this.mContext = context;
        this.walletList = walletListBeans;
    }


    public void setAddOnclickListener(onAddClickListener addClick) {
        this.addClickListener = addClick;
    }


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.dia_wallect);
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
        mImg_close = findViewById(R.id.img_close);
        mImg_add2 = findViewById(R.id.img_add2);
        mRecycler_wallect = findViewById(R.id.recycler_wallect);
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
        RecyclerView.LayoutManager layoutManager = new LinearLayoutManager(mContext);
        mRecycler_wallect.setLayoutManager(layoutManager);
        WallectListAdapter mainBtcAdapter = new WallectListAdapter(walletList);
        mRecycler_wallect.setAdapter(mainBtcAdapter);


    }

    /**
     * 初始化界面的确定和取消监听
     */
    private void initEvent() {
        mImg_close.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                dismiss();
//                }
            }
        });

        //设置取消按钮被点击后，向外界提供监听
        mImg_add2.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
//                if (noOnclickListener != null) {
//                    noOnclickListener.onNoClick();
//                }
            }
        });
    }

    /**
     * 从外界Activity为Dialog设置标题
     *
     * @param title
     */
    public void setTitle(String title) {
        titleStr = title;
    }

    /**
     * 从外界Activity为Dialog设置message
     *
     * @param message
     */
    public void setMessage(String message) {
        messageStr = message;
    }

    public interface onAddClickListener {
        public void onAddClick();
    }


}
