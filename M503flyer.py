import time
from opensky_api import OpenSkyApi

def m2ft(altitude):
    if altitude:
        altitude = round(float(altitude) * 3.2808399, 1)
    return altitude

def precheck(ac):
    if ac.latitude is None or ac.latitude > 28 or ac.latitude < 18:
        return False
    if ac.longitude is None or ac.longitude > 129 or ac.longitude < 114:
        return False
    if ac.callsign.strip() == '':
        return False

    return True

def M503check(ac):
    # BEGMO = ('BEGMO', 27.59, 121.50)
    OKATO = ('OKATO', 27.35, 121.34)
    NUDPO = ('NUDPO', 26.46, 121.04)
    PONEN = ('PONEN', 25.38, 120.23)
    OBKEL = ('OBKEL', 24.59, 119.52)
    APAKA = ('APAKA', 23.51, 118.26)
    # TOLAK = ('TOLAK', 23.06, 117.29)
    M503 = [OKATO, NUDPO, PONEN, OBKEL, APAKA]

    for w in M503:
        dist = (ac.latitude - w[1])*(ac.latitude - w[1]) + \
            (ac.longitude - w[2])*(ac.longitude - w[2])
        if dist < 0.1: #make it small enough to near waypoint
            print(ac.callsign, w[0], ac.heading, round(dist, 3))

def postcheck():
    pass

def main():
    api = OpenSkyApi()
    s = api.get_states()

    flights = []
    for ac in s.states:
        status = None
        ac.geo_altitude = m2ft(ac.geo_altitude)

        if not precheck(ac):
            continue

        M503check(ac)
    #     if M503check(ac):
    #         status = (ac.callsign.strip(), ac.latitude, ac.longitude, ac.geo_altitude)
        
    #     flights.append(status)

    # with open('ac', 'w') as ac:
    #     for f in flights:
    #         ac.writelines(str(f)+'\n')

if __name__ == '__main__':
    while True:
        main()
        print('-------------------')
        time.sleep(5)
        
    
