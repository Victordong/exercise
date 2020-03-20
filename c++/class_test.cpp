#include <string.h>
#include <sys/epoll.h>
#include <sys/time.h>
#include <sys/timerfd.h>
#include <unistd.h>
#include <condition_variable>
#include <iostream>
#include <map>
#include <mutex>
#include <vector>

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
    Result() { std::cout << "Result()" << std::endl; };
    Result(const Result& r) { std::cout << "&Result()" << std::endl; };
    ~Result() { std::cout << "~Result()" << std::endl; };
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
    std::cout << epfd << " " << timerfd << std::endl;
    struct epoll_event ev;
    ev.data.fd = timerfd;
    ev.events = EPOLLIN;
    int result = epoll_ctl(epfd, EPOLL_CTL_ADD, timerfd, &ev);
    std::cout << result << std::endl;
    itimerspec howlong;
    bzero(&howlong, sizeof howlong);
    howlong.it_value.tv_sec = 5;
    result = ::timerfd_settime(timerfd, 0, &howlong, NULL);
    std::cout << result << std::endl;
    while (true) {
        std::cout << "heihei" << std::endl;
        int numEvents =
            epoll_wait(epfd, &*events.begin(), events.size(), 1000 * 10);
        if (numEvents != 0) {
            std::cout << "123" << std::endl;
        }
    }
    ::close(epfd);
    ::close(timerfd);
}

Result ProcessShape(const Shape& first, const Shape& second) {
    return Result();
}

int main(int argc, char const* argv[]) {
    Result&& a = ProcessShape(Circle(), Triangle());
    return 0;
}
