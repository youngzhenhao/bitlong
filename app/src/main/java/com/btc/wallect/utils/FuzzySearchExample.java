package com.btc.wallect.utils;

import com.btc.wallect.model.entity.AddMnemonBean;

import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

public class FuzzySearchExample {
    public static boolean isMatch(String query, String candidate) {
        query = query.toLowerCase(Locale.getDefault());
        candidate = candidate.toLowerCase(Locale.getDefault());
        int n = query.length();
        int m = candidate.length();
        int[][] dp = new int[n + 1][m + 1];

        for (int i = 0; i <= n; i++) {
            for (int j = 0; j <= m; j++) {
                if (i == 0) {
                    dp[i][j] = 0;
                } else if (j == 0) {
                    dp[i][j] = 0;
                } else {
                    if (query.charAt(i - 1) == candidate.charAt(j - 1)) {
                        dp[i][j] = dp[i - 1][j - 1] + 1;
                    } else {
                        dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
                    }
                }
            }
        }
        return dp[n][m] >= 1;
    }

    public static List<AddMnemonBean> queryList(List<AddMnemonBean> mqueryList,String queryTxt) {


        ArrayList<AddMnemonBean> matchedItems = new ArrayList<>();
        ArrayList<AddMnemonBean> matchedItemsR = new ArrayList<>();
        for (AddMnemonBean item : mqueryList) {
            if (isMatch(queryTxt, item.getCollect())) {
                matchedItems.add(item);
            }
        }

        // 输出匹配的结果
        for (AddMnemonBean match : matchedItems) {
            System.out.println(match);
            AddMnemonBean addMnemonBean=new AddMnemonBean();
            addMnemonBean.setState(match.isState());
            addMnemonBean.setCollect(match.getCollect());
            addMnemonBean.setIndex(match.getIndex());
            matchedItemsR.add(addMnemonBean);
        }
        // Log.e("lz>>>liasst:",new Gson().toJson(matchedItemsR));
        return matchedItemsR;
    }
}
