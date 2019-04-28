package Singleton;

public class EagerSingleton {
    private static EagerSingleton instance = new EagerSingleton();
    private EagerSingleton (){}

    /**
     * 线程安全：是
     * 难度：易
     * 描述：最常用的实现方式，但容易产生垃圾对象。
     *      它基于 classloader 机制避免了多线程的同步问题，不过，instance 在类装载时就实例化，虽然导致类装载的原因有很多种，在单例模式中大多数都是调用 getInstance 方法， 但是也不能确定有其他的方式（或者其他的静态方法）导致类装载，这时候初始化 instance 显然没有达到 lazy loading 的效果。
     * @return
     */
    public static EagerSingleton getInstance() {
        return instance;
    }
}
