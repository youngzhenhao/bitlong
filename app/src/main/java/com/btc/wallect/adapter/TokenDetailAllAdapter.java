package com.btc.wallect.adapter;

import android.graphics.Color;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.btc.wallect.model.entity.WalletListBean;
import com.btc.wallect.R;

import java.util.List;

public class TokenDetailAllAdapter extends RecyclerView.Adapter<TokenDetailAllAdapter.ViewHolder> {
private List<WalletListBean> mFruitList;//用于存放数据的实体类数组


static class ViewHolder extends RecyclerView.ViewHolder {
    TextView btcBtcKey;
    TextView item_tv_btc_time;
    TextView mItem_tv_amonut;
    ImageView mImagCopy,mtem_img;


    public ViewHolder(@NonNull View itemView) {
        super(itemView);
        btcBtcKey = itemView.findViewById(R.id.item_tv_btcKey);
        item_tv_btc_time = itemView.findViewById((R.id.item_tv_btc_time));
        mItem_tv_amonut = itemView.findViewById((R.id.item_tv_amonut));
        mImagCopy=itemView.findViewById(R.id.item_img_copy);
        mtem_img=itemView.findViewById(R.id.item_img);
    }

}

    public TokenDetailAllAdapter(List<WalletListBean> fruitList) {
        mFruitList = fruitList;
    }


    @NonNull
    @Override
    public TokenDetailAllAdapter.ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_token_alldetail, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        TokenDetailAllAdapter.ViewHolder viewHolder = new TokenDetailAllAdapter.ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }


    public void onBindViewHolder(TokenDetailAllAdapter.ViewHolder viewHolder, int i) {
        WalletListBean fruit = mFruitList.get(i);      //获取实体类数组中数据
        //将数据bind到子项中控件（子项控件已缓存到了ViewHolder了）
        viewHolder.btcBtcKey.setText(fruit.getWallectKey()+"");
        viewHolder.mItem_tv_amonut.setText(fruit.getWallectAmount()+"");
        if(fruit.getType()==001){
            viewHolder.mtem_img.setImageResource(R.mipmap.img_tureafer);
            viewHolder.mItem_tv_amonut.setTextColor(Color.parseColor("#50AF95"));
        }else {
            viewHolder.mtem_img.setImageResource(R.mipmap.img_btc_in);
            viewHolder.mItem_tv_amonut.setTextColor(Color.parseColor("#2A82E4"));
        }

    }

    //作用：返回RecyclerView子项总数
    @Override
    public int getItemCount() {
        return mFruitList.size();
    }
}
