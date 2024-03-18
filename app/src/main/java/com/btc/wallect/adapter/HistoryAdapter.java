package com.btc.wallect.adapter;

import android.graphics.Color;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.R;
import com.btc.wallect.model.entity.HistoryBean;
import com.btc.wallect.utils.DateTimeUtil;


import java.util.List;

public class HistoryAdapter extends RecyclerView.Adapter<RecyclerView.ViewHolder> {

    private static final int TYPE_HEADER = 0;
    private static final int TYPE_ITEM = 1;
    private List<HistoryBean> items;
    private View headerView;

    public HistoryAdapter(List<HistoryBean> items, View headerView) {
        this.items = items;
        this.headerView = headerView;
    }

    @Override
    public RecyclerView.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        if (viewType == TYPE_HEADER) {
            return new HeaderViewHolder(headerView);
        } else {
            View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_history_nom, parent, false);
            return new ItemViewHolder(view);
        }
    }

    @Override
    public void onBindViewHolder(RecyclerView.ViewHolder holder, int position) {
        if (holder instanceof ItemViewHolder) {
            // 绑定数据到正常的item视图
            HistoryBean item = items.get(position - 1); // 减去头部视图所占的位置
            ((ItemViewHolder) holder).bindData(item);
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
        TextView mItemtvAmount;
        TextView mItem_tv_state;
        TextView mItem_tv_data;
        TextView mItem_tv_time;
        public ItemViewHolder(View itemView) {
            super(itemView);
            mItemtvAmount=itemView.findViewById(R.id.item_tv_amount);
            mItem_tv_state=itemView.findViewById(R.id.item_tv_state);
            mItem_tv_data=itemView.findViewById(R.id.item_tv_data);
            mItem_tv_time=itemView.findViewById(R.id.item_tv_time);
        }

        public void bindData(HistoryBean data) {
            // 绑定数据到item视图的控件
            mItemtvAmount.setText(""+data.getAmount()+"sats");

            if(data.getState()==0){
                mItem_tv_state.setTextColor(Color.parseColor("#EC3468"));
                mItem_tv_state.setText(R.string.app_txt57);
            }else if (data.getState()==1){
                mItem_tv_state.setTextColor(Color.parseColor("#2A82E4"));
                mItem_tv_state.setText(R.string.app_txt58);
            }else {
                mItem_tv_state.setTextColor(Color.parseColor("#38D176"));
                mItem_tv_state.setText(R.string.app_txt59);
            }
            mItem_tv_data.setText(DateTimeUtil.getInstance().formatDateTime(data.getDataTime(),14));
            mItem_tv_time.setText(DateTimeUtil.getInstance().formatDateTime(data.getDataTime(),11));
        }
    }
}
