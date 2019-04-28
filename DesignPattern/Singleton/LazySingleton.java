package Singleton;

public class LazySingleton {
    private static LazySingleton instance;
    private LazySingleton (){}

    /**
     * 线程安全：否
     * 难度：易
     * 描述：最基本的实现方式，但是因为没有加锁 synchronized，所以严格意义上它并不算单例模式。
     */
    public static LazySingleton getInstance() {
        if (instance == null) {
            instance = new LazySingleton();
        }
        return instance;
    }

    /**
     * 线程安全：是
     * 难度：易
     * 描述：具备很好的 lazy loading，能够在多线程中很好的工作，但是效率很低，99% 情况下不需要同步。
     */
    public static synchronized LazySingleton getInstanceWithSynchronized() {
        if (instance == null) {
            instance = new LazySingleton();
        }
        return instance;
    }

    /**
     * 线程安全：是
     * 难度：较复杂
     * 描述：这种方式采用双锁机制，安全且在多线程情况下能保持高性能。
     */
    private volatile static LazySingleton singleton;
    public static LazySingleton getSingleton() {
        if (singleton == null) {
            synchronized (LazySingleton.class) {
                if (singleton == null) {
                    singleton = new LazySingleton();
                }
            }
        }
        return singleton;
    }

    /**
     * 线程安全：是
     * 难度：一般
     * 描述：这种方式能达到双检锁方式一样的功效，但实现更简单。对静态域使用延迟初始化，应使用这种方式而不是双检锁方式。这种方式只适用于静态域的情况，双检锁方式可在实例域需要延迟初始化时使用。
     *      这种方式同样利用了 classloader 机制来保证初始化 instance 时只有一个线程。
     *      饿汉式只要 Singleton 类被装载了，那么 instance 就会被实例化（没有达到 lazy loading 效果）
     *      这种方式是 Singleton 类被装载了，instance 不一定被初始化。因为 SingletonHolder 类没有被主动使用，只有通过显式调用 getInstance 方法时，才会显式装载 SingletonHolder 类，从而实例化 instance。
     */
    private static class SingletonHolder {
        private static final LazySingleton INSTANCE = new LazySingleton();
    }
    public static final LazySingleton getInstanceUseSingletonHolder() {
        return SingletonHolder.INSTANCE;
    }
}
