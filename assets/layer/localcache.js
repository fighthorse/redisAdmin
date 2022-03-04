/**
 * **********************************************************************
 * 说明：    缓存对象LRU策略
 * 声明方式：
 *    indexKey 缓存名
 *    cacheSize 缓存个数
 *
 *   var localCache = new LRUCache(indexKey, cacheSize);
 * 调用方式：
 *   添加：localCache.put(key, value, expires);
 *   获取： localCache.get(key);
 *   清除：localCache.clear();
 *   删除：localCache.remove(key);
 *   删除最后一个：localCache.removeLast();
 * **********************************************************************
 */
/**
 * 缓存对象
 * @param indexKey 缓存名
 * @param cacheSize 缓存个
 * @constructor
 */
function LRUCache(indexKey, cacheSze) {
    // 秒 分 时 一天
    var expiresTime = 60 * 24 * 365;
    var currentSize = 0; // 当前大小
    var cacheSize = cacheSze; //缓存大小
    var keyArray = [];//缓存容器
    var localStorage = window.localStorage || {
        getItem: function () {
            return null;
        },
        setItem: function () {
        }
    }; // 缓存全局localStorage对象局部变量
    indexKey = sha1(indexKey);
    /**
     * 加载localStorage中的数据,并且判断是否有忆失效的
     */
    keyArray = JSON.parse(getLocalStorage(indexKey)) || [];
    currentSize = keyArray.length;
    for (var i = 0; i < keyArray.length; i++) {
        var time = +keyArray[i].time;
        var nowTime = +new Date();
        if (nowTime > time) {
            removeLocalStorage(keyArray[i].key);
            keyArray.splice(i, 1);
        }
        saveKeyIndexLocalStorage();
    }

    /**
     * hash加密->sha1算法
     * @param str
     */
    function sha1(str) {
        return CryptoJS.SHA1(str).toString();
    }

    /**
     * 获取字符串长度（汉字算两个字符，字母数字算一个）
     * @author @DYL
     * @DATE 2017-11-21
     */
    function getByteLen(val) {
        var len = 0;
        for (var i = 0; i < val.length; i++) {
            var a = val.charAt(i);
            if (a.match(/[^\x00-\xff]/ig) != null) {
                len += 2;
            } else {
                len += 1;
            }
        }
        return len;
    }

    function saveKeyIndexLocalStorage(obj) {
        if (obj) {
            keyArray.unshift(obj);
        }
        localStorage.setItem(indexKey, JSON.stringify(keyArray));
    }

    /**
     * 把数据保存到localStorage中
     * @expire 单独是分钟
     * @param obj Object类型
     */
    function saveLocalStorage(key, value) {
        localStorage.setItem(key, value);
    }

    function removeLocalStorage(key) {
        removeLastLocalStorage(key)
    }

    function removeLastLocalStorage(key) {
        currentSize--;
        localStorage.removeItem(key);
    }

    function getLocalStorage(key) {
        return localStorage.getItem(key);
    }

    /**
     * 添加缓存
     * @param key
     * @param value
     */
    function put(key, value, expires) {
        // 判断是不是已经存在了缓存对象
        var item = get(key);
        value = JSON.stringify(value); //将对象转换成字符JSON
        key = sha1(key);
        if (item) {
            moveToHead(key);
        } else {
            //缓存容器是否已经超过大小.
            if (currentSize >= cacheSize) {
                removeLast();
            } else {
                currentSize++;
                expires = expires || expiresTime;
                expires = +new Date() + expires * 60 * 1000; // 默认一天失效
                saveKeyIndexLocalStorage({key: key, time: expires, size: getByteLen(value)});
            }
        }
        saveLocalStorage(key, value);
    }

    /**
     * 获取缓存中对象
     * @param key
     * @return
     */
    function get(key) {
        key = sha1(key);
        moveToHead(key);
        return JSON.parse(getLocalStorage(key));
    }

    /**
     * 将缓存删除
     * @param key
     * @return
     */
    function remove(key) {
        key = sha1(key);
        for (var i = 0, len = keyArray.length; i < len; i++) {
            if (key == keyArray[i].key) {
                currentSize--;
                keyArray.splice(i, 1);
                removeLocalStorage(key);
                saveKeyIndexLocalStorage();
                break;
            }
        }
    }

    /**
     * 清仓缓存cachaName中的数据
     */
    function clear() {
        keyArray = [];
        currentSize = 0;
        localStorage.clear();
    }

    /**
     * 删除链表尾部节点
     *  表示 删除最少使用的缓存对象
     */
    function removeLast() {
        var key = keyArray.pop();
        removeLastLocalStorage(key);
    }

    /**
     * 移动到链表头，表示这个节点是最新使用过的
     * @param node
     */
    function moveToHead(key) {
        var item = {};
        for (var i = 0; i < keyArray.length; i++) {
            item = keyArray[i];
            if (key == item.key && i != 0) {
                //如果不是第一个元素，就把它从数字中删除，再把它添加到数组顶端
                item = keyArray.splice(i, 1)[0];
                saveKeyIndexLocalStorage(item);
                break;
            }
        }
    }

    this.get = get;
    this.put = put;
    this.remove = remove;
    this.clear = clear;
    this.removeLast = removeLast;
    this.moveToHead = moveToHead;
}

window.LRUCache = LRUCache;