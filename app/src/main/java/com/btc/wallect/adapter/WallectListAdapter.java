package com.btc.wallect.adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.R;
import com.btc.wallect.model.Imoder.onItemClickListener;
import com.btc.wallect.model.Imoder.onListItemListener;
import com.btc.wallect.model.entity.Wallet;
import com.btc.wallect.model.entity.WalletListBean;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.CopyUtil;

import java.util.List;

public class WallectListAdapter extends RecyclerView.Adapter<WallectListAdapter.ViewHolder> {
private List<Wallet> mFruitList;//用于存放数据的实体类数组
    public onListItemListener listener;
    public void setonItemClickListener(onListItemListener listener) {
        this.listener = listener;
    }


static class ViewHolder extends RecyclerView.ViewHolder {
    TextView btcName;
    TextView item_tv_btc;
    TextView mItem_tv_amonut;
    ImageView mImagCopy;


    public ViewHolder(@NonNull View itemView) {
        super(itemView);
        btcName = itemView.findViewById(R.id.item_tv_btc_name);
        item_tv_btc = itemView.findViewById((R.id.item_tv_btc));
        mItem_tv_amonut = itemView.findViewById((R.id.item_tv_amonut));
        mImagCopy=itemView.findViewById(R.id.item_img_copy);
    }

}

    public WallectListAdapter(List<Wallet> fruitList) {
        mFruitList = fruitList;
    }


    @NonNull
    @Override
    public WallectListAdapter.ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_wallect_list, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        WallectListAdapter.ViewHolder viewHolder = new WallectListAdapter.ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }


    public void onBindViewHolder(WallectListAdapter.ViewHolder viewHolder, int i) {
        Wallet fruit = mFruitList.get(i);      //获取实体类数组中数据
        //将数据bind到子项中控件（子项控件已缓存到了ViewHolder了）
        viewHolder.btcName.setText(fruit.name);
        viewHolder.item_tv_btc.setText(fruit.btcKey+"");
        viewHolder.mItem_tv_amonut.setText(fruit.btcAmount+"BTC");
        viewHolder.mImagCopy.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                CopyUtil.copyClicks(fruit.btcKey);
            }
        });
        viewHolder.itemView.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if(listener != null) {
                    int position = viewHolder.getAdapterPosition();
                    if(position != RecyclerView.NO_POSITION)
                        listener.onItemClick(fruit.id);
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
