#include <string.h>
#include <sys/timerfd.h>
#include <memory>
#include "archer/net/eventloop/channel.hpp"

using namespace archer;

archer::Eventloop* g_loop;

void timeout() {
    std::cout << "Time out" << std::endl;
    g_loop->quit();
}

using Entry = std::pair<std::string, int>;
using EntryList = std::vector<Entry>;
using Map = std::map<string, EntryList>;
using MapPtr = std::shared_ptr<Map>;

Maptr CustomData::getData() {
    std::lock_guard<std::mutex>(lock_);
    return data_;
}

int CustomData::query(const std::string& customer, const std::string& stock) {
    MapPtr data = getData();
    Map::const_iterator entries = data->find(customer);
    if (entries != data->end()) {
        return findEntry(entries->second, stock);
    } else {
        return -1;
    }
}

void CustomData::update(const std::string& customer, const EntryList& entries) {
    std::lock_guard<std::mutex>(lock_);
    if (!data_.unique()) {
        MapPtr newData(new Map(*data_));
    }
    (*data_)[customer] = entries;
}

MapPtr parseData(const std::string& message);
void CustomData::update(const std::string& message) {
    MapPtr newData = parseData(message);
    if (newData) {
        std::lock_guard<std::mutex>(lock_);
        data_.swap(newData);
    }
}

int main(int argc, char const* argv[]) {
    archer::Eventloop loop;

    g_loop = &loop;

    int timerfd = timerfd_create(CLOCK_MONOTONIC, TFD_NONBLOCK | TFD_CLOEXEC);
    archer::Channel channel(&loop, timerfd);
    loop.AddChannel(channel);
    channel.set_read_callback(timeout);
    channel.EnableReading();

    itimerspec howlong;
    bzero(&howlong, sizeof howlong);
    howlong.it_value.tv_sec = 1;
    int result = timerfd_settime(timerfd, 0, &howlong, nullptr);

    loop.Loop();

    close(timerfd);
    std::cout << "end" << std::endl;
    return 0;
}
