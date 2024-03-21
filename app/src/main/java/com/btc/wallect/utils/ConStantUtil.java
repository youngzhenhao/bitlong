package com.btc.wallect.utils;


import android.os.Environment;

import java.io.File;

public class ConStantUtil {
    public static String FILE_PAGE="file_page";
    public static String ACC="123456";
    public static String TRUE="true";
    public static String FALSE="false";
    public static String KEY_TOACTION="UI_page";
    public static String V_TOACTION_CREATE="UI_page_create";
    public static String V_TOACTION_INPUT="UI_page_input";
    public static  String ADD_MNEMON="add_mnemon_word";
    public static  String ADD_MNEMON_EDIT="add_edit_mnemon_word";

    public static String MAIN_TAB_LIST="main_tab_list";
    public static String ISWALLECT="is_wallect";
    public static String CURRENT_SQL_ID="current_sql_id";

    public static boolean STATE_TRUE=true;
    public static boolean STATE_FALSE=false;
    public static String WALLECT_STATE="wallect_state";
    public static String WALLECT_VERIFY="wallect_verify";

    public static String WALLECT_CREATE="wallect_create";
    public static String WALLECT_IMPORT="wallect_import";
    public static String WALLECT_ACCOUT="wallect_account";
    public static String WALLECT_UNZIP="wallect_unzip";


    public static String ConfTxt() {
        String line="\r\n";
        String t1 = "alias=Andy1 LND1 Taproot" + line;
        String t2 = "[Bitcoin]" + line;
        String t3 = "bitcoin.active=true" + line;
        String t4 = "bitcoin.node=bitcoind" + line;
        String t5 = "bitcoin.testnet=true" + line+line;

        String t6 = "[Bitcoind]" + line;
        String t7 = "bitcoind.rpchost=144.91.118.23" + line;
        String t8 = "bitcoind.dir=~/.bitcoin" + line;
        String t9 = "bitcoind.rpcuser=user" + line;
        String t10 = "bitcoind.rpcpass=password" + line;
        String t11 = "bitcoind.zmqpubrawblock=tcp://144.91.118.23:28332" + line;
        String t12 = "bitcoind.zmqpubrawtx=tcp://144.91.118.23:28333" + line;
        String t13 = "[rpcmiddleware]" + line;
        String t14 = "rpcmiddleware.enable=true" + line;
        String t15 = "[bolt]" + line;
        String t16 = "db.bolt.auto-compact=true" + line;
        return t1 + t2 + t3 + t4 + t5 + t6 + t7 + t8 + t9 + t10 + t11 + t12 + t13 + t14 + t15 + t16;
    }
    public static String ConfigTxt() {
        String line="\r\n";
        String t1 = line+"lndhost=127.0.0.1:8990" + line;
        String t2 = "taproothost=127.0.0.1:8991" + line;

        return t1 + t2;
    }
    public static String getPathFile(){
        File externalStorage = Environment.getExternalStorageDirectory();
        String path = externalStorage.getAbsolutePath();
        return path;
    }


}
