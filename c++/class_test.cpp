#include <string.h>
#include <sys/epoll.h>
#include <sys/time.h>
#include <sys/timerfd.h>
#include <unistd.h>
#include <condition_variable>
#include <functional>
#include <iostream>
#include <map>
#include <mutex>
#include <set>
#include <vector>

#include <stdint.h>

class noncopyable {
   protected:
    noncopyable() = default;
    ~noncopyable() = default;
    noncopyable(const noncopyable&) = delete;
    noncopyable& operator=(const noncopyable&) = delete;
};

class Shape {
   public:
    Shape() { std::cout << "Shape()" << std::endl; };
    virtual ~Shape() { std::cout << "~Shape()" << std::endl; };
    void Print() { std::cout << "Print" << std::endl; };
    void set_callback(const std::function<void()>& cb) { callback_ = cb; };
    void callback() { callback_(); };

   private:
    std::function<void()> callback_;
};

class Circle : public Shape {
   public:
    Circle() { std::cout << "Circle()" << std::endl; };
    Circle(Circle&& s) { std::cout << "&&Circle()" << std::endl; };
    ~Circle() { std::cout << "~Circle()" << std::endl; };
};

class Triangle : public Shape {
   public:
    Triangle() { std::cout << "Triangle()" << std::endl; };
    Triangle(Triangle&& t) { std::cout << "&&Triangle()" << std::endl; };
    ~Triangle() { std::cout << "~Triangle()" << std::endl; };
};

class Result {
   public:
    Result() : num_(0) { std::cout << "Result()" << std::endl; };
    Result(int num) : num_(num) { std::cout << "Result()" << std::endl; };
    Result(const Result& r) : num_(r.num_) {
        std::cout << "&Result()" << std::endl;
    };
    ~Result() { std::cout << "~Result()" << std::endl; };
    void Print() { std::cout << "Print" << std::endl; };
    void Delete() { num_--; };
    void UniquePtrTest(const std::shared_ptr<Result>& r,
                       std::vector<std::shared_ptr<Result>>& list);
    int num() { return num_; };

    void set_num(int num) { num_ = num; };

   private:
    int num_;
};

class SomeTest : noncopyable {
   public:
    SomeTest() = default;
    SomeTest(std::string& s) : str(s){};
    SomeTest(const char* s) : str(s){};
    ~SomeTest() = default;
    SomeTest& operator=(const SomeTest& s) {
        str = s.str;
        return *this;
    };

   private:
    std::string str;
};

class World {
   private:
    /* data */

    std::mutex lock;
    std::condition_variable condition;

   public:
    World(/* args */);
    ~World();

    void Test() {
        using std::swap;
        lock.lock();
        lock.unlock();
    };
};

World::World(/* args */) {}

World::~World() {}

void test_epoll() {
    std::vector<epoll_event> events(10);
    int epfd = epoll_create1(EPOLL_CLOEXEC);
    int timerfd = timerfd_create(CLOCK_MONOTONIC, TFD_NONBLOCK | TFD_CLOEXEC);
    struct epoll_event ev;
    ev.data.fd = timerfd;
    ev.events = EPOLLIN;
    int result = epoll_ctl(epfd, EPOLL_CTL_ADD, timerfd, &ev);
    itimerspec howlong;
    bzero(&howlong, sizeof howlong);
    howlong.it_value.tv_sec = 5;
    result = ::timerfd_settime(timerfd, 0, &howlong, NULL);
    while (true) {
        int numEvents =
            epoll_wait(epfd, &*events.begin(), events.size(), 1000 * 10);
        if (numEvents != 0) {
            uint64_t buf;
            int num = ::read(timerfd, &buf, sizeof(uint64_t));

            result = ::timerfd_settime(timerfd, 0, &howlong, NULL);
            num = ::read(timerfd, &buf, sizeof(uint64_t));
        }
    }
    ::close(epfd);
    ::close(timerfd);
}

Result ProcessShape(const Shape& first, const Shape& second) {
    return Result();
}

void Result::UniquePtrTest(const std::shared_ptr<Result>& r,
                           std::vector<std::shared_ptr<Result>>& list) {
    r->Print();
    r->Delete();
    list.push_back(r);
    std::cout << &list << std::endl;

    std::cout << "remain " << r.use_count() << std::endl;
}

class Timestamp final {
   public:
    Timestamp() : micro_seconds_(0){};

    explicit Timestamp(int ms) : micro_seconds_(ms){};

    Timestamp(const Timestamp& ts) : micro_seconds_(ts.micro_seconds_){};

    ~Timestamp() = default;

    static const int kMicroSecondsPerSecond = 1000 * 1000;

    static const int kNanoSecondsPerSecond = 1000 * 1000 * 1000;

    static const int kNanoSecondsPerMicroSecond = 1000;

    bool operator<(const Timestamp& other) const {
        return micro_seconds_ < other.micro_seconds_;
    };

    bool operator>(const Timestamp& other) const {
        return micro_seconds_ > other.micro_seconds_;
    };

    bool operator==(const Timestamp& other) const {
        return micro_seconds_ == other.micro_seconds_;
    };

    Timestamp& operator+(int delay) {
        micro_seconds_ += micro_seconds_ + delay;
        return *this;
    }

    int expiration() {
        struct timeval tv;
        gettimeofday(&tv, nullptr);
        int ms = int(tv.tv_sec) * kMicroSecondsPerSecond + int(tv.tv_usec);
        return micro_seconds_ - ms;
    };

    void swap(Timestamp& first, Timestamp& second) {
        using std::swap;
        swap(first.micro_seconds_, second.micro_seconds_);
    };

    static Timestamp Now() {
        struct timeval tv;
        gettimeofday(&tv, nullptr);
        int ms = int(tv.tv_sec) * kMicroSecondsPerSecond + int(tv.tv_usec);
        return Timestamp(ms);
    };

    int micro_seconds() const { return micro_seconds_; };

   private:
    int micro_seconds_;
};

void map_test() {
    std::map<int, int> m;
}

void test_share_ptr() {
    std::shared_ptr<Shape> s(new Shape());
    std::vector<std::shared_ptr<Result>> list;

    {
        std::shared_ptr<Result> r(new Result());
        std::cout << r << std::endl;
        std::weak_ptr<Result> wr(r);
        std::cout << "before insert " << r.use_count() << " weak_ptr "
                  << wr.use_count() << std::endl;
        list.push_back(r);
        std::cout << "after insert " << r.use_count() << " weak_ptr "
                  << wr.use_count() << std::endl;
        s->set_callback([&]() {
            std::cout << " before Print " << r.use_count() << " weak_ptr "
                      << wr.use_count() << std::endl;
            r->Print();
            std::cout << r << std::endl;
            std::cout << "after Print " << r.use_count() << " weak_ptr "
                      << wr.use_count() << std::endl;
            list.clear();
            std::cout << "after empty " << r.use_count() << " weak_ptr "
                      << wr.use_count() << std::endl;
            r->Print();
            r->set_num(11);
            std::cout << r << std::endl;

            std::cout << r->num() << std::endl;
        });
        std::cout << "after set callback " << r.use_count() << " weak_ptr "
                  << wr.use_count() << std::endl;
    }

    std::cout << list[0] << std::endl;

    s->callback();
}

int main() {
    std::map<int, int> m;
    std::cout << "123" << std::endl;
    return 0;
}