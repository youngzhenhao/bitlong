package com.btc.wallect.adapter;

import android.app.Activity;
import android.content.res.Resources;
import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.R;
import com.btc.wallect.model.entity.HistoryBean;
import com.btc.wallect.model.entity.TaprTabListBean;
import com.btc.wallect.utils.DateTimeUtil;

import java.util.List;

public class TaprTabAdapter extends RecyclerView.Adapter<RecyclerView.ViewHolder> {

    private static final int TYPE_HEADER = 0;
    private static final int TYPE_ITEM = 1;
    private List<TaprTabListBean> items;
    private View headerView;
    public static Activity activity;

    public TaprTabAdapter(Activity mActivity,List<TaprTabListBean> items, View headerView) {
        this.activity=mActivity;
        this.items = items;
        this.headerView = headerView;
    }

    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        if (viewType == TYPE_HEADER) {
            return new TaprTabAdapter.HeaderViewHolder(headerView);
        } else {
            View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_tapr_tab, parent, false);
            return new TaprTabAdapter.ItemViewHolder(view);
        }
    }

    @Override
    public void onBindViewHolder(RecyclerView.ViewHolder holder, int position) {
        if (holder instanceof TaprTabAdapter.ItemViewHolder) {
            // 绑定数据到正常的item视图
            TaprTabListBean item = items.get(position - 1); // 减去头部视图所占的位置
            ((TaprTabAdapter.ItemViewHolder) holder).bindData(item);
        }
    }

    @Override
    public int getItemCount() {
        return items.size() + 1; // 加上头部视图的数量
    }

    @Override
    public int getItemViewType(int position) {
        if (position == 0) {
            return TYPE_HEADER;
        }
        return TYPE_ITEM;
    }

    public static class HeaderViewHolder extends RecyclerView.ViewHolder {

        public HeaderViewHolder(View itemView) {
            super(itemView);
            // 初始化头部视图的控件
        }
    }

    public static class ItemViewHolder extends RecyclerView.ViewHolder {
        TextView mTv_1,mTv_2,mTv_3,mTv_4,mTv_5;

        public ItemViewHolder(View itemView) {
            super(itemView);
            mTv_1=itemView.findViewById(R.id.tv_1);
            mTv_2=itemView.findViewById(R.id.tv_2);
            mTv_3=itemView.findViewById(R.id.tv_3);
            mTv_4=itemView.findViewById(R.id.tv_4);
            mTv_5=itemView.findViewById(R.id.tv_5);
        }

        public void bindData(TaprTabListBean data) {
            // 绑定数据到item视图的控件
            mTv_1.setText(""+data.getDealName());
            mTv_2.setText("$"+data.getDealAmount()+"BTC");
            mTv_3.setText("$"+data.getDealdDetailPrice()+"stas");

            if(data.getState()==0){
                mTv_5.setText("+"+data.getRose());
                Resources resources = activity.getResources();
                Drawable btnDrawable = resources.getDrawable(R.drawable.bg_green);
                mTv_5.setBackground(btnDrawable);
            }else if (data.getState()==1){
                Resources resources = activity.getResources();
                Drawable btnDrawable2 = resources.getDrawable(R.drawable.bg_red);
                mTv_5.setBackground(btnDrawable2);
                mTv_5.setText("-"+data.getRose());
            }
        }
    }
}
