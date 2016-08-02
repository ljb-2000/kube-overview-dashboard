$(function() {

    Morris.Area({
        element: 'morris-pods-by-nodes',
        data: [{
            period: '2010 Q1',
            master1: 2,
            node1: null,
            node2: 1
        }, {
            period: '2010 Q2',
            master1: 2,
            node1: 1,
            node2: null
        }, {
            period: '2010 Q3',
            master1: 2,
            node1: 3,
            node2: 5
        }, {
            period: '2010 Q4',
            master1: 2,
            node1: 2,
            node2: 1
        }, {
            period: '2011 Q1',
            master1: 3,
            node1: 6,
            node2: 7
        }],
        xkey: 'period',
        ykeys: ['master1', 'node1', 'node2'],
        labels: ['Master1 ', 'Node 1', 'Node 2'],
        pointSize: 2,
        hideHover: 'auto',
        resize: true
    });

    Morris.Donut({
        element: 'morris-donut-chart',
        data: [{
            label: "Download Sales",
            value: 12
        }, {
            label: "In-Store Sales",
            value: 30
        }, {
            label: "Mail-Order Sales",
            value: 20
        }],
        resize: true
    });

    Morris.Bar({
        element: 'morris-bar-chart',
        data: [{
            y: '2006',
            a: 100,
            b: 90
        }, {
            y: '2007',
            a: 75,
            b: 65
        }, {
            y: '2008',
            a: 50,
            b: 40
        }, {
            y: '2009',
            a: 75,
            b: 65
        }, {
            y: '2010',
            a: 50,
            b: 40
        }, {
            y: '2011',
            a: 75,
            b: 65
        }, {
            y: '2012',
            a: 100,
            b: 90
        }],
        xkey: 'y',
        ykeys: ['a', 'b'],
        labels: ['Series A', 'Series B'],
        hideHover: 'auto',
        resize: true
    });

});
