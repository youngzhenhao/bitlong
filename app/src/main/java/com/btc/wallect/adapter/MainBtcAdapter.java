package com.btc.wallect.adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.model.Imoder.onItemClickListener;
import com.btc.wallect.R;
import com.btc.wallect.model.entity.MainDateilBean;

import java.util.List;

public class MainBtcAdapter extends RecyclerView.Adapter<MainBtcAdapter.ViewHolder> {
    private List<MainDateilBean> mFruitList;//用于存放数据的实体类数组

    public onItemClickListener listener;

    public void setonItemClickListener(onItemClickListener listener) {
        this.listener = listener;
    }
    static class ViewHolder extends RecyclerView.ViewHolder {
        TextView mItemTvBtc;
        TextView mItemTvMode;
        TextView mItemTvAmonut;
        TextView mItemTvAll;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);
            mItemTvBtc = itemView.findViewById(R.id.item_tv_btc);
            mItemTvMode = itemView.findViewById((R.id.item_tv_mode));
            mItemTvAmonut = itemView.findViewById((R.id.item_tv_amonut));
            mItemTvAll = itemView.findViewById((R.id.item_tv_all));
        }

    }

    public MainBtcAdapter(List<MainDateilBean> fruitList) {
        mFruitList = fruitList;
    }


    @NonNull
    @Override
    public MainBtcAdapter.ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_main_btc, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        MainBtcAdapter.ViewHolder viewHolder = new MainBtcAdapter.ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }


    public void onBindViewHolder(MainBtcAdapter.ViewHolder viewHolder, int i) {
        MainDateilBean fruit = mFruitList.get(i);      //获取实体类数组中数据
        //将数据bind到子项中控件（子项控件已缓存到了ViewHolder了）
        viewHolder.mItemTvBtc.setText(fruit.getBtcName());
        viewHolder.mItemTvMode.setText(fruit.getBtcMode());
        viewHolder.mItemTvAmonut.setText(fruit.getBtcAmount()+"");
        viewHolder.mItemTvAll.setText("$"+String.valueOf(fruit.getBtcAll()));
        viewHolder.itemView.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (listener != null) {
                    listener.onItemClick();
                }
            }
        });
    }

    //作用：返回RecyclerView子项总数
    @Override
    public int getItemCount() {
        return mFruitList.size();
    }
}
