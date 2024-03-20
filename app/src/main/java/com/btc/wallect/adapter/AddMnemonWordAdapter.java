package com.btc.wallect.adapter;


import android.graphics.Color;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;


import com.btc.wallect.R;
import com.btc.wallect.db.DBOpenHelper;
import com.btc.wallect.db.DBdao;
import com.btc.wallect.model.Imoder.onItemClickListener;
import com.btc.wallect.model.entity.AddMnemonBean;
import com.btc.wallect.model.entity.CollectBean;
import com.btc.wallect.model.entity.Wallet;

import java.util.List;

public class AddMnemonWordAdapter extends RecyclerView.Adapter<AddMnemonWordAdapter.ViewHolder> {
    private List<AddMnemonBean> mFruitList;

    private List<CollectBean> sqlCollectList;
    public onItemClickListener listener;

    public void setonItemClickListener(onItemClickListener listener) {
        this.listener = listener;
    }

    static class ViewHolder extends RecyclerView.ViewHolder {
        TextView tv_index;
        TextView tv_collect;

        public ViewHolder(@NonNull View itemView) {
            super(itemView);

            tv_index = itemView.findViewById((R.id.tv_txt_index));
            tv_collect = itemView.findViewById((R.id.tv_add_txt));
        }
    }

    public AddMnemonWordAdapter(List<AddMnemonBean> fruitList, List<CollectBean> walletList) {
        mFruitList = fruitList;
        sqlCollectList = walletList;

    }


    @NonNull
    @Override
    public AddMnemonWordAdapter.ViewHolder onCreateViewHolder(@NonNull ViewGroup viewGroup, int i) {
        //动态加载布局
        View view = LayoutInflater.from(viewGroup.getContext()).inflate(R.layout.item_add_mnemon_word, viewGroup, false);
        //创建ViewHolder实例，参数为刚加载进来的子项布局
        AddMnemonWordAdapter.ViewHolder viewHolder = new AddMnemonWordAdapter.ViewHolder(view);//这样，子项布局里面的控件就缓存到了ViewHolder

        return viewHolder;
    }

    public void onBindViewHolder(AddMnemonWordAdapter.ViewHolder viewHolder, int i) {
        AddMnemonBean fruit = mFruitList.get(i);      //获取实体类数组中数据
        viewHolder.tv_index.setText(fruit.getIndex() + "");
        if(!fruit.isState()){
            viewHolder.tv_collect.setTextColor(Color.parseColor("#EC3468"));
        }else {
            viewHolder.tv_collect.setTextColor(Color.parseColor("#383838"));
        }
//        if (sqlCollectList.size() > 0) {
//            if (!fruit.getCollect().isEmpty() && !fruit.getCollect().equals("")) {
//                if (!sqlCollectList.get(i).getName().equals(fruit.getCollect())) {
//                    viewHolder.tv_collect.setTextColor(Color.parseColor("#EC3468"));
//                }
//            }else {
//                viewHolder.tv_collect.setTextColor(Color.parseColor("#383838"));
//            }
//        }

        viewHolder.tv_collect.setText(fruit.getCollect());

    }


    @Override
    public int getItemCount() {
        return mFruitList.size();
    }

}
